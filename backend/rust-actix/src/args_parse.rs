use clap::Parser;

/// Simple todo api
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub struct Args {
    /// port to run the server on
    #[arg(short, long, default_value_t = 4000)]
    pub port: u16,
}
