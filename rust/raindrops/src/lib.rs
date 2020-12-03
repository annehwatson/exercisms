pub fn raindrops(n: u32) -> String {
    if n < 3 {
        return n.to_string();
    }
    let mut result = String::new();
    if n % 3 == 0 {
        result.push_str("Pling");
    }
    if n % 5 == 0 {
        result.push_str("Plang");
    }
    if n % 7 == 0 {
        result.push_str("Plong");
    }
    if result.len() == 0 {
        return n.to_string();
    } else {
        return result.to_string();
    }
}
