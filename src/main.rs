use std::fs;
use std::env;
use std::path::PathBuf;
use std::io;

//TODO: Need to validate if file is valid yaml
fn load_yaml(file_name: &str) -> Result<String, io::Error> {
    let root_dir = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
    let file_path = root_dir.join(file_name);

    let contents = fs::read_to_string(file_path)?;

    return Ok(contents);
}


fn main() -> std::io::Result<()> {
    let file_name = "test.yaml";
    let contents = load_yaml(file_name)?;
    println!("contents: {contents}");
    return Ok(());
}
