extern crate project_euler;

use project_euler::solver;

fn main() {
    println!("Hello, project euler!");
    println!("{}", solver::number_spiral_diagonals_028::solver());
}
