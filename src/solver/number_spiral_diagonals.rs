// problem no.28

const SIZE: u64 = (1001 - 1) / 2;

fn sum_center_right_up() -> u64 {
    (1..SIZE).map(|e| e * e).fold(0, |acc, x| acc + x)
}

pub fn solver() -> u64 {
    sum_center_right_up()
}
