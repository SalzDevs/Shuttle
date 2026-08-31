use std::fs;
use std::env;
use std::path::PathBuf;

fn main() -> std::io::Result<()> {
    let root_dir = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
    
    let file_path = root_dir.join("test.yaml");
    println!("{}", file_path.display());

    let contents = fs::read_to_string(file_path)?;
    
    println!("{contents}");

    return Ok(());
}
