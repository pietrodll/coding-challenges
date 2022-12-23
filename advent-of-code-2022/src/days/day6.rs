use std::collections::HashSet;

fn find_first_distint_index(string: &str, size: usize) -> usize {
    let mut last_chars: HashSet<char> = HashSet::new();

    for idx in 0..(string.len() - size) {
        last_chars.extend(string[idx..(idx + size)].chars());

        if last_chars.len() >= size {
            return idx + size;
        }

        last_chars.drain();
    }

    return string.len();
}

pub fn run_first_part(data: &String) -> String {
    find_first_distint_index(data, 4).to_string()
}

pub fn run_second_part(data: &String) -> String {
    find_first_distint_index(data, 14).to_string()
}
