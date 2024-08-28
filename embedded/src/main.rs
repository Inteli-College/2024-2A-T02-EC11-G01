use std::io::{self, Write};
use std::process::Command;
use chrono;

use v4l::buffer::Type;
use v4l::io::traits::CaptureStream;
use v4l::prelude::*;
use v4l::video::Capture;

fn main() -> io::Result<()> {
    let path = "/dev/video0";
    println!("Using device: {}\n", path);

    let count = 50; // Number of frames to capture
    let buffer_count = 4; // Number of buffers

    let dev = Device::with_path(path)?;
    let mut format = dev.format()?;
    format.fourcc = v4l::FourCC::new(b"YUYV"); // Explicitly set YUYV format
    dev.set_format(&format)?;

    let params = dev.params()?;
    println!("Active format:\n{}", format);
    println!("Active parameters:\n{}", params);

    let mut stream = MmapStream::with_buffers(&dev, Type::VideoCapture, buffer_count)?;

    // Prepare the output path in the home directory
    let current_dir = std::env::current_dir().unwrap();
    let output_path = format!("{}/images/{}.mp4", current_dir.display(), chrono::prelude::Utc::now());

    // Start an ffmpeg process to encode the video stream to MP4
    let mut ffmpeg = Command::new("ffmpeg")
        .arg("-f")
        .arg("rawvideo") // Specify that the input is raw video
        .arg("-pix_fmt")
        .arg("yuyv422") // Assume YUYV format
        .arg("-s") // Frame size (width x height)
        .arg(format!("{}x{}", format.width, format.height))
        .arg("-r")
        .arg("30") // Frame rate (30 fps)
        .arg("-i")
        .arg("-") // Input from stdin (raw video data)
        .arg("-c:v")
        .arg("libx264") // Use H.264 encoder
        .arg("-preset")
        .arg("fast")
        .arg("-y") // Overwrite output file
        .arg(&output_path) // Output path
        .stdin(std::process::Stdio::piped()) // Pipe stdin for raw data
        .spawn()
        .expect("Failed to start ffmpeg");

    let stdin = ffmpeg.stdin.as_mut().expect("Failed to open stdin");

    // Capture and send frames to ffmpeg
    for _ in 0..count {
        let (buf, meta) = stream.next()?;

        // Ensure buffer size matches expected frame size
        if buf.len() != (format.width * format.height * 2) as usize {
            eprintln!(
                "Skipping frame due to incorrect buffer size: {} (expected: {})",
                buf.len(),
                (format.width * format.height * 2)
            );
            continue;
        }

        stdin.write_all(&buf).expect("Failed to write to ffmpeg");
        println!(
            "Captured frame with sequence: {}, timestamp: {}",
            meta.sequence, meta.timestamp
        );
    }

    // Close stdin to let ffmpeg finish encoding
    drop(stdin);

    // Wait for ffmpeg to finish
    ffmpeg.wait().expect("ffmpeg process failed");

    println!("Video saved to {}", output_path);

    Ok(())
}
