pub fn find_fuel_requirement(mass: i32) -> i32 {
    mass / 3 - 2
}

pub fn find_fuel_fuel_requirement(mass: i32) -> i32 {
    let mut sum: i32 = mass;
    let mut last: i32 = mass;

    loop {
        let current: i32 = find_fuel_requirement(last);
        if current <= 0 {
            break
        }
        sum = sum + current;
        last = current;
    }

    sum
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_find_fuel_requirement() {
        assert_eq!(find_fuel_requirement(12), 2);
        assert_eq!(find_fuel_requirement(14), 2);
        assert_eq!(find_fuel_requirement(1969), 654);
        assert_eq!(find_fuel_requirement(100756), 33583);
    }

    #[test]
    fn test_find_fuel_fuel_requirement() {
        assert_eq!(find_fuel_fuel_requirement(2), 2);
        assert_eq!(find_fuel_fuel_requirement(654), 966);
        assert_eq!(find_fuel_fuel_requirement(33583), 50346);
    }
}
