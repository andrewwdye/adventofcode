use clap::Parser;
use nalgebra as na;
use std::fs::read_to_string;


#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String
}

fn main() -> Result<(), std::io::Error>{
    let cli = Cli::parse();
    let input = read_to_string(cli.input)?;
    let result = match cli.part {
        1 => solve1(input.as_str())?,
        2 => solve2(input.as_str())?,
        _ => unreachable!(),
    };
    println!("{result}");
    Ok(())
}

struct Machine {
    a: (i32, i32),
    b: (i32, i32),
    prize: (i32, i32)
}

impl Machine {
    fn new(lines: &mut std::str::Lines) -> Self {
        let re = regex::Regex::new(r"\d+").unwrap();
        let a: Vec<i32> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        let b: Vec<i32> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        let prize: Vec<i32> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        _ = lines.next();
        Machine { a: (a[0], a[1]), b: (b[0], b[1]), prize: (prize[0], prize[1]) }
    }
}

fn solve1(input: &str) -> Result<i64, std::io::Error>{
    let mut lines = input.lines();
    let mut machines = Vec::new();
    while lines.clone().peekable().peek().is_some() {
        let m = Machine::new(&mut lines);
        machines.push(m);
    }
    let mut sum = 0;
    for m in machines {
        for a_presses in 0..=100 {
            let ax = m.a.0 * a_presses;
            let ay = m.a.1 * a_presses;
            let bx = m.prize.0 - ax;
            if bx % m.b.0 != 0 {
                continue;
            }
            let b_presses = bx / m.b.0;
            if ay + m.b.1 * b_presses != m.prize.1 {
                continue;
            }
            let tokens = 3 * a_presses + b_presses;
            sum += tokens;
        }
    }
    Ok(sum as i64)
}

struct Machine2 {
    a: (i64, i64),
    b: (i64, i64),
    prize: (i64, i64)
}

impl Machine2 {
    fn new(lines: &mut std::str::Lines) -> Self {
        let re = regex::Regex::new(r"\d+").unwrap();
        let a: Vec<i64> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i64>().unwrap()).collect();
        let b: Vec<i64> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i64>().unwrap()).collect();
        let prize: Vec<i64> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i64>().unwrap()).collect();
        _ = lines.next();
        Machine2 { a: (a[0], a[1]), b: (b[0], b[1]), prize: (prize[0] + 10000000000000, prize[1] + 10000000000000) }
    }
}

fn solve2(input: &str) -> Result<i64, std::io::Error>{
    let mut lines = input.lines();
    let mut machines = Vec::new();
    while lines.clone().peekable().peek().is_some() {
        let m = Machine2::new(&mut lines);
        machines.push(m);
    }
    let mut sum = 0;
    for m in machines {
        let counts = na::Matrix2::new(m.a.0 as f64, m.b.0 as f64, m.a.1 as f64, m.b.1 as f64);
        let prize = na::Vector2::new(m.prize.0 as f64, m.prize.1 as f64);
        let lu: nalgebra::LU<f64, nalgebra::Const<2>, nalgebra::Const<2>> = na::LU::new(counts);
        let presses = lu.solve(&prize).unwrap();
        let a_presses = presses[0].round() as i64;
        let b_presses = presses[1].round() as i64;
        if m.a.0 * a_presses + m.b.0 * b_presses == m.prize.0 && m.a.1 * a_presses + m.b.1 * b_presses == m.prize.1 {
            let tokens = 3 * a_presses + b_presses;
            sum += tokens;
        }
    }
    Ok(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 480);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT).unwrap(), 875318608908);
    }
}
