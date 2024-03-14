
import java.util.*;

public class References {

    public static void main(String[] args) {
//          Increment hashmap value by 1 if it exists, if not, set it to 1
        HashMap<Integer, Integer> map = new HashMap<>();
        Integer key = 0;

        map.merge(key, 1, Integer::sum);


//          Sort based on hashmap values
        //List of the keys
        List<Integer> l = new ArrayList<>(map.keySet());
        //Sort the list based on hashmap value
        Collections.sort(l, new Comparator<Integer>() {
            public int compare(Integer a, Integer b) {
                return (map.get(b) - map.get(a));
            }
        });



        //Get hashmap key with largest value
//        Key key = Collections.max(map.entrySet(), Map.Entry.comparingByValue()).getKey();

        //Max value in hashmap
        int maxValueInMap = (Collections.max(map.values()));


//          Sort based on string length. Lambda shortened expression of comparator
//        Collections.sort(words, (a, b) -> {
//            if (a.length() != b.length()) {
//                return a.length() - b.length(); // Sort by length
//            }
//            return a.compareTo(b); // If lengths are the same, sort lexicographically
//        });


//        Char to int
        String n = "123";
        int i = 0;
        int x = n.charAt(i) - '0';


//        Map char to an int
//        a = 0, b = 1,...,z = 25
//        A = -32, B = -31,...,Z = -7
        n = "abc";
        x = n.charAt(0) - 'a';
        System.out.println(x);


        //Int to binary
//        Integer.toBinaryString(int i)

//        PriorityQueue
//        Insertion order is not retained, the order of the queue is based on the specified comparator
//        From smallest to biggest:
        PriorityQueue<Integer> PQ = new PriorityQueue<>();
//        From biggest to smallest:
        PriorityQueue<Integer> PQ2 = new PriorityQueue<>(Collections.reverseOrder());

        PriorityQueue<Integer> queue = new PriorityQueue<>((a,b) -> (a-b));


//        Condensed for loop for strings
        String s = "abcdefg";
        for (char c : s.toCharArray()) {
        }


//        Double int arrays
        int[] a = {1, 2, 3};
        int[] b = {4, 5, 6};
        int[][] doubleArray = {a, b};
//        doubleArray[0][0] will be 1
//        doubleArray[1][0] will be 4
//        doubleArray[y][x]
//        doubleArray.length will be 2

        //Sub array
        int[] subArray = Arrays.copyOfRange(a, 0, 0);


//        Instantiating List using Arrays.asList()
        List<Integer> list = Arrays.asList(1, 2, 3);



        //Int array to list
        List<Integer> listt = Arrays.stream(a).boxed().toList();

        //List to int array
//        int[] output = list.stream().mapToInt(i->i).toArray();


        //Sort by reverse order
        Collections.sort(listt, Collections.reverseOrder());

        //Sort a set
//        Set<Integer> keySet = treeMapKey.keySet();
//        List<Integer> listSet = new ArrayList<>(keySet);
//        Collections.sort(listSet, Collections.reverseOrder());


//        Finding the middle of a linkedlist, use tortoise and hare method.
    }

}

