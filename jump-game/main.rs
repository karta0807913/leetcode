pub struct Solution {}

impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut index: usize = 0;
        let mut max_index: usize = 0;
        while index <= max_index && index < nums.len() {
            max_index = std::cmp::max(max_index, index + nums[index] as usize);
            if max_index >= (nums.len() - 1) {
                return true;
            }
            index += 1;
        }
        return false;
    }
}

#[cfg(test)]
mod tests {
    use crate::jump_game::Solution;

    #[test]
    fn hello_na() {
        print!("hi");
        println!("HELLLLLLO");
    }

    #[test]
    fn test() {
        assert_eq!(
            Solution::can_jump(vec![2,3,1,1,4]),
            true,
        );
        assert_eq!(
            Solution::can_jump(vec![3,2,1,0,4]),
            false,
        );
        assert_eq!(
            Solution::can_jump(vec![2,0]),
            true,
        );
        assert_eq!(
            Solution::can_jump(vec![2,5,0,0]),
            true,
        );
        assert_eq!(
            Solution::can_jump(vec![0]),
            true,
        );
        assert_eq!(
            Solution::can_jump(vec![2,0,0]),
            true,
        );
    }
}
