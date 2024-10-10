// main.rs

mod features;

use std::fs::File;
use std::io::{self, Write};
use chrono;
use dotenv::dotenv;
use v4l::buffer::Type;
use v4l::io::traits::CaptureStream;
use v4l::prelude::*;
use v4l::video::Capture;
use image::{RgbImage, ImageBuffer};
use serialport::{self, DataBits, Parity, StopBits, FlowControl};
use std::time::Duration;
use byteorder::{ByteOrder, LittleEndian};

const DISTANCE_THRESHOLD: f32 = 40.0; // Define the distance threshold in centimeters

#[tokio::main]
async fn main() -> io::Result<()> {
    dotenv().ok();
    
    let path = "/dev/video0";
    println!("Using device: {}\n", path);

    let buffer_count = 4;

    // Open the device
    let dev = Device::with_path(path)?;
    
    // Set the format (YUYV)
    let mut format = dev.format()?;
    format.fourcc = v4l::FourCC::new(b"YUYV"); // YUYV format
    dev.set_format(&format)?;

    let params = dev.params()?;
    println!("Active format:\n{}", format);
    println!("Active parameters:\n{}", params);

    // Start capturing
    let mut stream = MmapStream::with_buffers(&dev, Type::VideoCapture, buffer_count)?;

    let port_name = "/dev/ttyACM0";
    let mut port = serialport::new(port_name, 9600) // Set baud rate
        .data_bits(DataBits::Eight) // Set data bits
        .parity(Parity::None)
        .stop_bits(StopBits::One)
        .flow_control(FlowControl::None)
        .timeout(Duration::from_millis(1000))
        .open()?;

    let mut serial_buf: Vec<u8> = vec![0; 4]; // Buffer for float data

    loop {
        // Read data from serial port
        match port.read(serial_buf.as_mut_slice()) {
            Ok(bytes_read) if bytes_read == 4 => { // Ensure 4 bytes are read

                let mut distance = LittleEndian::read_f32(&serial_buf);
                
                // Clamp distance to range [0, 30]
                if distance > DISTANCE_THRESHOLD || distance < 0.0 {
                    distance = 0.0;
                }
                
                println!("Processed distance: {}", distance);
                
                // Check if the distance exceeds the threshold
                if distance > 30.0 {
                    // Capture a single frame
                    let (buf, meta) = stream.next()?;
                    
                    // Ensure buffer size matches expected frame size
                    if buf.len() != (format.width * format.height * 2) as usize {
                        eprintln!(
                            "Skipping frame due to incorrect buffer size: {} (expected: {})",
                            buf.len(),
                            (format.width * format.height * 2)
                        );
                    } else {
                        
                        let rgb_data = yuyv_to_rgb(&buf, format.width as usize, format.height as usize);

                        
                        let current_dir = std::env::current_dir().unwrap();
                        let timestamp = chrono::prelude::Utc::now().format("%Y%m%d_%H%M%S").to_string();
                        let output_path = format!("{}/images/{}.png", current_dir.display(), timestamp);

                        
                        let img: RgbImage = ImageBuffer::from_raw(format.width, format.height, rgb_data)
                            .expect("Failed to create image buffer");
                        img.save(&output_path).expect("Failed to save image");

                        println!(
                            "Captured frame with sequence: {}, timestamp: {}",
                            meta.sequence, meta.timestamp
                        );

                        println!("Image saved to {}", output_path);

                        if let Err(e) = features::aws::s3_upload::file_to_upload(&output_path).await {
                            eprintln!("Error uploading file: {:?}", e);
                        }
                    }
                }
            }
            Ok(bytes_read) => {
                println!("");
            }
            Err(e) => eprintln!("Error reading from serial port: {:?}", e),
        }
    }
}


fn yuyv_to_rgb(yuyv: &[u8], width: usize, height: usize) -> Vec<u8> {
    let mut rgb = Vec::with_capacity(width * height * 3);

    for chunk in yuyv.chunks(4) {
        let y1 = chunk[0] as f32;
        let u = chunk[1] as f32;
        let y2 = chunk[2] as f32;
        let v = chunk[3] as f32;

        rgb.extend_from_slice(&yuv_to_rgb(y1, u, v));
        rgb.extend_from_slice(&yuv_to_rgb(y2, u, v));
    }

    rgb
}

fn yuv_to_rgb(y: f32, u: f32, v: f32) -> [u8; 3] {
    let c = y - 16.0;
    let d = u - 128.0;
    let e = v - 128.0;

    let r = (1.164 * c + 1.596 * e).clamp(0.0, 255.0) as u8;
    let g = (1.164 * c - 0.392 * d - 0.813 * e).clamp(0.0, 255.0) as u8;
    let b = (1.164 * c + 2.017 * d).clamp(0.0, 255.0) as u8;

    [r, g, b]
}
