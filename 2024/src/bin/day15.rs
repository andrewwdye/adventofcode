use clap::Parser;
use std::fs::read_to_string;

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String,
}

fn main() -> Result<(), std::io::Error> {
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

#[derive(PartialEq, Eq, Clone, Copy)]
enum Direction {
    Up = 0,
    Down,
    Left,
    Right,
}

static STEPS: &'static [(i32, i32)] = &[(0, -1), (0, 1), (-1, 0), (1, 0)];

impl Direction {
    fn parse(lines: &mut std::str::Lines) -> Vec<Direction> {
        let mut dirs = Vec::new();
        for line in lines {
            for c in line.chars() {
                let dir = match c {
                    '^' => Direction::Up,
                    'v' => Direction::Down,
                    '<' => Direction::Left,
                    '>' => Direction::Right,
                    _ => panic!("Invalid character in directions"),
                };
                dirs.push(dir);
            }
        }
        dirs
    }
}

impl std::fmt::Debug for Direction {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Direction::Up => write!(f, "^"),
            Direction::Down => write!(f, "v"),
            Direction::Left => write!(f, "<"),
            Direction::Right => write!(f, ">"),
        }
    }
}

#[derive(PartialEq, Eq, Clone, Copy, Debug)]
enum Content {
    Wall,
    Empty,
    Box,
    LeftBox,
    RightBox,
    Robot,
}

struct Warehouse {
    grid: Vec<Vec<Content>>,
    robot: (i32, i32),
}

impl Warehouse {
    fn new(lines: &mut std::str::Lines) -> Self {
        let mut grid = Vec::new();
        let mut robot = (0, 0);
        loop {
            let line = lines.next().unwrap();
            if line.is_empty() {
                break;
            }
            let mut row = Vec::new();
            for (x, c) in line.chars().enumerate() {
                match c {
                    '#' => row.push(Content::Wall),
                    '.' => row.push(Content::Empty),
                    'O' => row.push(Content::Box),
                    '@' => {
                        row.push(Content::Robot);
                        robot = (x as i32, grid.len() as i32);
                    }
                    _ => panic!("Invalid character in warehouse"),
                }
            }
            grid.push(row);
        }
        Warehouse { grid, robot }
    }

    fn new2(lines: &mut std::str::Lines) -> Self {
        let mut grid = Vec::new();
        let mut robot = (0, 0);
        loop {
            let line = lines.next().unwrap();
            if line.is_empty() {
                break;
            }
            let mut row = Vec::new();
            for (x, c) in line.chars().enumerate() {
                match c {
                    '#' => row.append(&mut vec![Content::Wall, Content::Wall]),
                    '.' => row.append(&mut vec![Content::Empty, Content::Empty]),
                    'O' => row.append(&mut vec![Content::LeftBox, Content::RightBox]),
                    '@' => {
                        row.push(Content::Robot);
                        row.push(Content::Empty);
                        robot = (x as i32 * 2, grid.len() as i32);
                    }
                    _ => panic!("Invalid character in warehouse"),
                }
            }
            grid.push(row);
        }
        Warehouse { grid, robot }
    }

    fn move_robot(&mut self, dir: Direction) {
        // Iterate in direction until finding empty space or edge
        let (dx, dy) = STEPS[dir.clone() as usize];
        let (mut x, mut y) = self.robot;
        let (mut final_x, mut final_y) = (x, y);
        loop {
            let (nx, ny) = (x + dx, y + dy);
            if nx < 0 || nx >= self.grid[0].len() as i32 || ny < 0 || ny >= self.grid.len() as i32 {
                break;
            }
            match self.grid[ny as usize][nx as usize] {
                Content::Wall => break,
                Content::Empty => {
                    (final_x, final_y) = (nx, ny);
                    break;
                }
                Content::Box => {
                    (x, y) = (nx, ny);
                }
                _ => unreachable!(),
            }
        }
        // println!("Move: {:?}, Robot: {:?}, Target: {:?} : {:?}", dir, self.robot, (x, y), self.grid[x as usize][y as usize]);
        if (final_x, final_y) == self.robot {
            return;
        }
        // Move boxes between robot and empty space
        let (nx, ny) = (self.robot.0 + dx, self.robot.1 + dy);
        if self.grid[ny as usize][nx as usize] == Content::Box {
            self.grid[ny as usize][nx as usize] = Content::Empty;
            self.grid[final_y as usize][final_x as usize] = Content::Box;
        }
        // Move robot
        self.grid[self.robot.1 as usize][self.robot.0 as usize] = Content::Empty;
        self.grid[ny as usize][nx as usize] = Content::Robot;
        self.robot = (nx, ny);
    }

    fn can_move(&self, x: i32, y: i32, dir: Direction) -> bool {
        let (dx, dy) = STEPS[dir.clone() as usize];
        let (nx, ny) = (x + dx, y + dy);
        if nx < 0 || nx >= self.grid[0].len() as i32 || ny < 0 || ny >= self.grid.len() as i32 {
            return false;
        }
        match self.grid[ny as usize][nx as usize] {
            Content::Wall => return false,
            Content::Empty => return true,
            Content::Box => {
                return self.can_move(nx, ny, dir);
            }
            Content::LeftBox => {
                if dir == Direction::Left || dir == Direction::Right {
                    return self.can_move(nx, ny, dir);
                } else {
                    return self.can_move(nx, ny, dir) && self.can_move(nx + 1, ny, dir);
                }
            }
            Content::RightBox => {
                if dir == Direction::Left || dir == Direction::Right {
                    return self.can_move(nx, ny, dir);
                } else {
                    return self.can_move(nx, ny, dir) && self.can_move(nx - 1, ny, dir);
                }
            }
            _ => unreachable!(),
        }
    }

    fn move_recurse(&mut self, x: i32, y: i32, dir: Direction) {
        let (dx, dy) = STEPS[dir.clone() as usize];
        let (nx, ny) = (x + dx, y + dy);
        if nx < 0 || nx >= self.grid[0].len() as i32 || ny < 0 || ny >= self.grid.len() as i32 {
            unreachable!();
        }
        let current = self.grid[y as usize][x as usize];
        let next = self.grid[ny as usize][nx as usize];
        match next {
            Content::Box => {
                self.move_recurse(nx, ny, dir);
            }
            Content::LeftBox => {
                if dir == Direction::Left || dir == Direction::Right {
                    self.move_recurse(nx, ny, dir);
                } else {
                    self.move_recurse(nx, ny, dir);
                    self.move_recurse(nx + 1, ny, dir);
                }
            }
            Content::RightBox => {
                if dir == Direction::Left || dir == Direction::Right {
                    self.move_recurse(nx, ny, dir);
                } else {
                    self.move_recurse(nx, ny, dir);
                    self.move_recurse(nx - 1, ny, dir);
                }
            }
            _ => {}
        }
        self.grid[ny as usize][nx as usize] = current;
        self.grid[y as usize][x as usize] = Content::Empty;
    }

    fn move_robot2(&mut self, dir: Direction) {
        // Iterate in direction until finding empty space or edge
        if !self.can_move(self.robot.0, self.robot.1, dir) {
            return;
        }
        // Move boxes between robot and empty space
        self.move_recurse(self.robot.0, self.robot.1, dir);
        // Update robot
        let (dx, dy) = STEPS[dir.clone() as usize];
        let (nx, ny) = (self.robot.0 + dx, self.robot.1 + dy);
        self.robot = (nx, ny);
    }

    fn sum_coordinates(&self) -> usize {
        let mut sum = 0;
        for y in 0..self.grid.len() {
            for x in 0..self.grid[y].len() {
                match self.grid[y][x] {
                    Content::Box | Content::LeftBox => {
                        sum += x + 100 * y;
                    }
                    _ => {}
                }
            }
        }
        sum
    }
}

impl std::fmt::Debug for Warehouse {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for row in &self.grid {
            for content in row {
                let c = match content {
                    Content::Wall => '#',
                    Content::Empty => '.',
                    Content::Box => 'O',
                    Content::LeftBox => '[',
                    Content::RightBox => ']',
                    Content::Robot => '@',
                };
                write!(f, "{}", c)?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn solve1(input: &str) -> Result<usize, std::io::Error> {
    let mut lines = input.lines();
    let mut warehouse = Warehouse::new(&mut lines);
    let directions = Direction::parse(&mut lines);
    for d in directions {
        warehouse.move_robot(d);
    }
    Ok(warehouse.sum_coordinates())
}

fn solve2(input: &str) -> Result<usize, std::io::Error> {
    let mut lines = input.lines();
    let mut warehouse = Warehouse::new2(&mut lines);
    let directions = Direction::parse(&mut lines);
    for d in directions {
        // println!("{:?}", warehouse);
        warehouse.move_robot2(d);
    }
    // println!("{:?}", warehouse);
    Ok(warehouse.sum_coordinates())
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT1: &str = "########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<";

    const SAMPLE_INPUT2: &str = "##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^";

    const SAMPLE_INPUT3: &str = "#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT1).unwrap(), 2028);
        assert_eq!(solve1(SAMPLE_INPUT2).unwrap(), 10092);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT3).unwrap(), 618);
        assert_eq!(solve2(SAMPLE_INPUT2).unwrap(), 9021);
    }
}
