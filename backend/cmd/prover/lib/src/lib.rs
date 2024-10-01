use std::ffi::{CString, CStr};
use std::os::raw::c_char;
use http_body_util::Empty;
use hyper::{body::Bytes, Request, StatusCode};
use hyper_util::rt::TokioIo;
use tlsn_core::proof::TlsProof;
use tlsn_examples::run_notary;
use tlsn_prover::tls::{state::Notarize, Prover, ProverConfig};
use tokio_util::compat::{FuturesAsyncReadCompatExt, TokioAsyncReadCompatExt};

const SERVER_DOMAIN: &str = "example.com";
const USER_AGENT: &str = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36";

#[no_mangle]
pub extern "C" fn notarize_request() -> *mut c_char {
    let result = tokio::runtime::Runtime::new()
        .unwrap()
        .block_on(execute());

    match result {
        Ok(proof) => {
            let json_string = serde_json::to_string(&proof).unwrap();
            let c_string = CString::new(json_string).unwrap();
            c_string.into_raw()
        }
        Err(e) => {
            let error_message = format!("Error during notarization: {:?}", e);
            let c_string = CString::new(error_message).unwrap();
            c_string.into_raw()
        }
    }
}

#[no_mangle]
pub extern "C" fn free_string(s: *mut c_char) {
    if s.is_null() {
        return;
    }
    unsafe {
        CString::from_raw(s);
    }
}

async fn execute() -> Result<TlsProof, Box<dyn std::error::Error>> {
    let (prover_socket, notary_socket) = tokio::io::duplex(1 << 16);

    // Start local notary service
    tokio::spawn(run_notary(notary_socket.compat()));

    // Prover configuration
    let config = ProverConfig::builder()
        .id("example")
        .server_dns(SERVER_DOMAIN)
        .build()
        .unwrap();

    // Create a Prover and set it up with the Notary
    let prover = Prover::new(config)
        .setup(prover_socket.compat())
        .await
        .unwrap();

    // Connect to the server via TCP
    let client_socket = tokio::net::TcpStream::connect((SERVER_DOMAIN, 443)).await?;

    // Connect the Prover to the server
    let (mpc_tls_connection, prover_fut) = prover.connect(client_socket.compat()).await?;

    let mpc_tls_connection = TokioIo::new(mpc_tls_connection.compat());

    // Run the Prover task
    let prover_task = tokio::spawn(prover_fut);

    // Connect the HTTP client to the MPC TLS connection
    let (mut request_sender, connection) =
        hyper::client::conn::http1::handshake(mpc_tls_connection).await?;

    tokio::spawn(connection);

    // Build the HTTP request
    let request = Request::builder()
        .uri("/")
        .header("Host", SERVER_DOMAIN)
        .header("Accept", "*/*")
        .header("Accept-Encoding", "identity")
        .header("Connection", "close")
        .header("User-Agent", USER_AGENT)
        .body(Empty::<Bytes>::new())
        .unwrap();

    // Send the request and receive the response
    let response = request_sender.send_request(request).await?;

    if response.status() != StatusCode::OK {
        return Err(format!("Received non-OK response: {}", response.status()).into());
    }

    // Wait for the Prover to finish
    let prover = prover_task.await.unwrap()?;

    // Prepare for notarization
    let prover = prover.start_notarize();

    // Generate the proof without redactions
    Ok(build_proof_without_redactions(prover).await)
}

async fn build_proof_without_redactions(mut prover: Prover<Notarize>) -> TlsProof {
    let sent_len = prover.sent_transcript().data().len();
    let recv_len = prover.recv_transcript().data().len();

    let builder = prover.commitment_builder();
    let sent_commitment = builder.commit_sent(&(0..sent_len)).unwrap();
    let recv_commitment = builder.commit_recv(&(0..recv_len)).unwrap();

    // Finalize, returning the notarized session
    let notarized_session = prover.finalize().await.unwrap();

    // Create a proof for all data committed in this session
    let mut proof_builder = notarized_session.data().build_substrings_proof();

    // Reveal all public ranges
    proof_builder.reveal_by_id(sent_commitment).unwrap();
    proof_builder.reveal_by_id(recv_commitment).unwrap();

    let substrings_proof = proof_builder.build().unwrap();

    TlsProof {
        session: notarized_session.session_proof(),
        substrings: substrings_proof,
    }
}
