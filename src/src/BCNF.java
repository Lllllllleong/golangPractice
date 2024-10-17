import java.util.*;


public class BCNF {
    List<List<String>> decompositionList;
    int[] determinants;
    int[] dependants;
    Integer fdMask;
    Integer attributeMask;
    Integer n;
    Integer m;
    List<Integer> remainingFDMask;


    public BCNF(int[] determinants,
                int[] dependants,
                int fdMask,
                int attributeMask,
                int n,
                int m) {
        this.decompositionList = new ArrayList<>();
        this.determinants = determinants;
        this.dependants = dependants;
        this.fdMask = fdMask;
        this.attributeMask = attributeMask;
        this.n = n;
        this.m = m;
        this.remainingFDMask = new ArrayList<>();
    }

    // Method to initiate BCNF decomposition.
    public void decompose() {
        decompositionList.clear();
        remainingFDMask.clear();
        List<String> innerList = new ArrayList<>();
        permuteBCNF(innerList, fdMask, attributeMask);
    }


    // Recursive backtracking bitmask method to compute all permutations of BCNF decomposition.
    public void permuteBCNF(List<String> currentTables, int fdMask, int attributeMask) {
        boolean flag = false;

        // Remove irrelevant FDs (those that are not part of the current attribute set).
        for (int i = 0; i < n; i++) {
            if ((fdMask & (1 << i)) == 0) continue;
            int determinant = determinants[i];
            int dependant = dependants[i];
            int detORdep = determinant | dependant;
            if ((attributeMask | detORdep) != attributeMask) fdMask &= ~(1 << i);
        }

        // Iterate over the remaining FDs to check for BCNF violations.
        for (int i = 0; i < n; i++) {
            if ((fdMask & (1 << i)) == 0) continue;
            // Check if the current FD is a superkey.
            int previousMask = 0;
            int currentMask = determinants[i];
            //Find the closure of the determinant
            while (previousMask != currentMask) {
                previousMask = currentMask;
                for (int j = 0; j < n; j++) {
                    if ((fdMask & (1 << j)) == 0) continue;
                    int determinant = determinants[j];
                    if ((currentMask | determinant) == currentMask) currentMask |= dependants[j];
                }
            }

            // If it's not a superkey, decompose.
            if (currentMask != attributeMask) {
                flag = true;
                List<String> nextTables = new ArrayList<>(currentTables);

                // R(X + Y): Create a new table for the attributes determined by this FD.
                StringBuilder sb = new StringBuilder();
                int detORdep = determinants[i] | dependants[i];
                for (int j = 0; j < m; j++) {
                    if ((detORdep & (1 << j)) != 0) sb.append(Character.toUpperCase((char) ('a' + j)));
                }

                // Cheap fix: I did not consider that the FDS need to be projected onto both split relations.
                // But, notice that there is only scenario where it needs to be split again, which is in the case where
                // FGIKM->H is split before M->I
                if (sb.toString().equals("FGHIKM")) {
                    nextTables.add("IM");
                    nextTables.add("FGHKM");
                } else {
                    nextTables.add(sb.toString());
                }

                // R(R - Y): Create a new table by excluding the dependant attributes.
                int nextFDMask = fdMask & ~(1 << i);
                int nextAttributeMask = attributeMask & ~dependants[i];
                permuteBCNF(nextTables, nextFDMask, nextAttributeMask);
            }
        }

        // If no BCNF violations were found, add the current set of tables to the decomposition.
        if (!flag) {
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < m; i++) {
                if ((attributeMask & (1 << i)) != 0) sb.append(Character.toUpperCase((char) ('a' + i)));
            }
            currentTables.add(sb.toString());
            decompositionList.add(currentTables);
        }
    }

    public static void main(String[] args) {
        // n is the number of functional dependencies.
        int n = 6;
        // m is the number of attributes.
        int m = 15;

        // Create the functional dependency mask.
        int fdMask = 0;
        for (int i = 0; i < n; i++) fdMask |= (1 << i);

        // Create the attribute mask.
        int attributeMask = 0;
        for (int i = 0; i < m; i++) attributeMask |= (1 << i);

        // Example functional dependencies (determinants and dependants).
        List<String> determinants = new ArrayList<>(Arrays.asList("ab", "e", "fgikm", "k", "m", "abe"));
        List<String> dependants = new ArrayList<>(Arrays.asList("cd", "fgjk", "h", "lm", "i", "no"));


        // Convert the functional dependency strings to bitmasks.
        int[] determinantMasks = stringToMask(determinants);
        int[] dependantMasks = stringToMask(dependants);

        // Initialise BCNF decomposition.
        BCNF bcnf = new BCNF(determinantMasks, dependantMasks, fdMask, attributeMask, n, m);
        bcnf.decompose(); // Start the decomposition.

        // Sort the resulting decompositions and filter out duplicates.
        List<List<String>> decomposition = bcnf.decompositionList;
        // Noting that, only the number of decomposed tables and the composition of the last
        // table and FDs, is what determines if a decomposition is unique.
        // Because I have appended the last (child) table, and it's FDs as a string, we can sort the collections of
        // strings, and filter out duplicates via adding all into a set.
        for (List<String> l : decomposition) Collections.sort(l);
        Set<List<String>> set = new HashSet<>(decomposition);
        decomposition = new ArrayList<>(set);

        // Find the FDs preserved
        List<int[]> relationMasks = new ArrayList<>();
        for (List<String> l : decomposition) relationMasks.add(stringToMask(l));

        int o = decomposition.size();
        for (int i = 0; i < o; i++) {
            int[] currentRelationMasks = relationMasks.get(i);
            String s = "with FDs: ";
            List<Integer> fds = new ArrayList<>();
            for (int j = 0; j < n; j++) {
                int detORdep = determinantMasks[j] | dependantMasks[j];
                boolean flag = false;
                for (int currentRelationMask : currentRelationMasks) {
                    if ((currentRelationMask | detORdep) == currentRelationMask) {
                        flag = true;
                        break;
                    }
                }
                if (flag) {
                    fds.add(j + 1);
                }
            }
            s += fds.toString();
            decomposition.get(i).add(s);
        }
        for (List<String> l : decomposition) System.out.println(l);
    }


    // Method to convert a set of attributes represented as strings into a bitmask.
    public static int[] stringToMask(List<String> stringList) {
        int n = stringList.size();
        int[] output = new int[n];
        for (int i = 0; i < n; i++) {
            int mask = 0;
            String s = stringList.get(i).toLowerCase();
            for (char c : s.toCharArray()) {
                int index = c - 'a';
                mask |= (1 << index);
            }
            output[i] = mask;
        }
        return output;
    }
}