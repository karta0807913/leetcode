struct Solution;

impl Solution {
    pub fn can_visit_all_rooms(mut rooms: Vec<Vec<i32>>) -> bool {
        let mut visited_rooms = vec![false; rooms.len()];
        let mut current_room = vec![0];
        let mut next_room: Vec<i32> = vec![];
        while current_room.len() != 0 {
            while current_room.len() != 0 {
                let room_index = current_room.pop().unwrap() as usize;
                if visited_rooms[room_index] {
                    continue;
                }
                next_room.extend(rooms[room_index].iter());
                visited_rooms[room_index] = true;
            }
            std::mem::swap(&mut current_room, &mut next_room);
            // println!("current_room: {:?}", current_room);
            // println!("next_room: {:?}", next_room);
            next_room.clear();
        }
        for &b in visited_rooms.iter() {
            // println!("{}", b);
            if !b {
                return false;
            }
        }
        return true;
    }
}

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn test() {
        assert_eq!(
            Solution::can_visit_all_rooms(
                vec![vec![1], vec![2], vec![3], vec![]],
            ),
            true,
        );
        assert_eq!(
            Solution::can_visit_all_rooms(
                vec![vec![1,3],vec![3,0,1],vec![2],vec![0]],
            ),
            false,
        );
    }
}
