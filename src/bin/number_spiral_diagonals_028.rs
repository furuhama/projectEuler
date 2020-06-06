// problem no.28

const SIZE: u64 = (1001 - 1) / 2;

// for test case
// const SIZE: u64 = 2;

fn sum_center_right_up(size: u64) -> u64 {
    (0..size + 1)
        .map(|e| 2 * e + 1)
        .map(|e| e * e)
        .fold(0, |acc, x| acc + x)
}

fn sum_right_down(size: u64) -> u64 {
    (1..size + 1)
        .map(|e| (2 * e - 1).pow(2) + 2 * e)
        .fold(0, |acc, x| acc + x)
}

fn sum_left_down(size: u64) -> u64 {
    (1..size + 1)
        .map(|e| (2 * e).pow(2) + 1)
        .fold(0, |acc, x| acc + x)
}

fn sum_left_up(size: u64) -> u64 {
    (1..size + 1)
        .map(|e| (2 * e).pow(2) + (2 * e) + 1)
        .fold(0, |acc, x| acc + x)
}

fn main() {
    let ans = sum_center_right_up(SIZE) + sum_right_down(SIZE) + sum_left_down(SIZE) + sum_left_up(SIZE);

    println!("{}", ans);
}
