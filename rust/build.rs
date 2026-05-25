use std::fs;
use std::io::Write;
use std::path::Path;

fn main() {
    let src = Path::new("src");
    let out_dir = std::env::var("OUT_DIR").unwrap();
    let dest = Path::new(&out_dir).join("problem_mods.rs");

    let mut mods: Vec<String> = fs::read_dir(src)
        .expect("failed to read src/")
        .filter_map(|e| e.ok())
        .filter_map(|e| {
            let name = e.file_name().to_string_lossy().to_string();
            name.strip_prefix("p_")
                .and_then(|n| n.strip_suffix(".rs"))
                .map(|stem| format!("p_{stem}"))
        })
        .collect();

    mods.sort();

    let src_abs = fs::canonicalize(src).expect("failed to canonicalize src/");
    let mut out = fs::File::create(&dest).expect("failed to create output file");
    for m in &mods {
        let path = src_abs.join(format!("{m}.rs"));
        writeln!(out, "#[path = \"{}\"]", path.display()).unwrap();
        writeln!(out, "pub mod {m};").unwrap();
    }

    println!("cargo:rerun-if-changed=src");
}
