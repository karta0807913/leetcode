struct Solution {}

impl Solution {
    pub fn detect_capital_use(word: String) -> bool {
        let (mut all_capital, mut all_small_case) = (true, true);
        let bytes = word.into_bytes();
        let mut first_capital = bytes[0] < 'a' as u8;
        return bytes
            .into_iter()
            .map(|c| c < 'a' as u8)
            .enumerate()
            .all(|(index, is_capital)| {
                all_capital &= is_capital;
                all_small_case &= !is_capital;
                first_capital &= !is_capital || index == 0;
                return all_capital || all_small_case || first_capital;
            });
    }
}
