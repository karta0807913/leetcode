struct Solution {}

use std::collections::HashSet;

impl Solution {
    pub fn word_pattern(pattern: String, s: String) -> bool {
        let mut exists_map: HashSet<&str> = HashSet::new();
        let mut cache = vec![""; 26];
        let mut remains = pattern.len();
        let pattern = pattern.into_bytes();
        for (index, substr) in s.split(' ').enumerate() {
            if index >= pattern.len() {
                return false;
            }
            let pattern_index = (pattern[index] - 'a' as u8) as usize;
            if cache[pattern_index].is_empty() {
                if exists_map.contains(substr) {
                    return false;
                }
                cache[pattern_index] = substr;
                exists_map.insert(substr);
            } else if cache[pattern_index] != substr {
                return false;
            }
            remains -= 1;
        }
        return remains == 0;
    }
}

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn word_pattern() {
        assert_eq!(
            false,
            Solution::word_pattern("abba".to_string(), "dog dog dog dog".to_string(),)
        );

        assert_eq!(
            false,
            Solution::word_pattern("a".to_string(), "dog dog dog dog".to_string(),)
        );
    }
}
