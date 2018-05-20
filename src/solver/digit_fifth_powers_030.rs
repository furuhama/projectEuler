// no.30

fn pow_fifth_by_digit(mut num: u32) -> u32 {
    let mut digits: Vec<u32> = Vec::new();

    if num >= 1000000 {
        panic!("An error occured. Maximum number is 6 * 9^5.");
    }

    if num >= 100000 {
        digits.push(num / 100000);
        num = num - (num / 100000) * 100000;
    }

    if num >= 10000 {
        digits.push(num / 10000);
        num = num - (num / 10000) * 10000;
    }

    if num >= 1000 {
        digits.push(num / 1000);
        num = num - (num / 1000) * 1000;
    }

    if num >= 100 {
        digits.push(num / 100);
        num = num - (num / 100) * 100;
    }

    if num >= 10 {
        digits.push(num / 10);
        num = num - (num / 10) * 10;
    }

    digits.push(num);

    digits.iter().fold(0, |acc, x| acc + x.pow(5))
}

fn answer_vec() -> Vec<u32> {
    let mut ans: Vec<u32> = Vec::new();

    let maximum: u32 = 9u32.pow(5) * 6;

    for i in 2..maximum+1 {
        if pow_fifth_by_digit(i) == i {
            println!("{}", i);
            ans.push(i);
        }
    }

    ans
}

pub fn solver() -> u32 {
    answer_vec().iter().fold(0, |acc, x| acc + x)
}

#[test]
fn to_digit_test() {
    assert_eq!(pow_fifth_by_digit(10000), 1);
    assert_eq!(pow_fifth_by_digit(12340), 1300);
    assert_eq!(pow_fifth_by_digit(54748), 54748);
    assert_eq!(pow_fifth_by_digit(92727), 92727);
    assert_eq!(pow_fifth_by_digit(93084), 93084);
}
