use clap::Parser;
use std::io::{BufRead, BufReader};
use std::fs::File;

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String
}

fn main() -> Result<(), std::io::Error>{
    let cli = Cli::parse();
    let result = match cli.part {
        1 => solve1(&mut BufReader::new(File::open(cli.input)?))?,
        2 => solve2(&mut BufReader::new(File::open(cli.input)?))?,
        _ => unreachable!(),
    };
    print!("{result}");
    Ok(())
}

fn solve1<R: BufRead>(input: &mut R) -> Result<i32, std::io::Error>{
    unimplemented!()
}

fn solve2<R: BufRead>(input: &mut R) -> Result<i32, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_sample() {
    }
}
