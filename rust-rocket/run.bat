REM Set current folder to use nightly Rust, for Rocket
rustup override set nightly

REM Get stuff
rustup update

REM Run it
cargo run