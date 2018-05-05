// no.29

fn _distinct_powers() -> u32 {
    0
}

// while thinking(incomplete size)
fn distinct_numbers(numbers: u32, size: u32) -> u32 {
    let distincts = match numbers {
        6 => 0,
        5 => 0,
        4 => size + (size / 2 + size / 3 - size / 6) + (size / 3 + size / 4 - size / 12) + size / 4,
        3 => size + (size / 2 + size / 3 - size / 6) + size / 3,
        2 => size + size / 2,
        1 => size,
        _ => panic!("unexpected argument! numbers should be less than 7."),
    };
    distincts
}

pub fn solver() -> u32 {
    distinct_numbers(4, 10)
}
