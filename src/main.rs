use std::fs;
use std::env;
use std::path::PathBuf;
use std::io;


fn load_yaml(path: &str) -> Result<String, io::Error> {
    let root_dir = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
    let file_path = root_dir.join(path);
    
    let contents = fs::read_to_string(file_path)?;

    return Ok(contents);
}

fn main() -> std::io::Result<()> {

    let contents = load_yaml("test.yaml")?;    
    println!("{contents}");

    return Ok(());
}
