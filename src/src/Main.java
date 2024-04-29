import java.util.*;


public class Main {


    static HashMap<Long, List<Long>> areaPriceMap;
    static TreeMap<Long, SquareFootage> areaStatisticsMap;

    public static long valuation(long reqArea, List<Long> area, List<Long> price) {
        // Write your code here
        int n = area.size();
        double output = 0d;
        if (n == 0) {
            output = 1000 * reqArea;
        } else if (n == 1) {
            output = price.get(0);
        } else {
            //Fill in the area/prices map
            areaPriceMap = new HashMap<>();
            areaStatisticsMap = new TreeMap<>();
            for (int i = 0; i < n; i++) {
                long currentArea = area.get(i);
                long currentPrice = price.get(i);
                if (!areaPriceMap.containsKey(currentArea)) areaPriceMap.put(currentArea, new ArrayList<>());
                areaPriceMap.get(currentArea).add(currentPrice);
            }
            //Construct a SquareFootage object for every area size in areaPriceMap
            //Also filter outliers
            for (Long areaKey : areaPriceMap.keySet()) {
                SquareFootage currentSF = new SquareFootage(areaPriceMap.get(areaKey));
                currentSF = currentSF.outlierFilter();
                if (currentSF != null) areaStatisticsMap.put(areaKey, currentSF);
            }
            //Valuation
            if (areaStatisticsMap.containsKey(reqArea)) {
                //If a reference home exists
                SquareFootage sf = areaStatisticsMap.get(reqArea);
                output = sf.meanPrice;
            } else {
                //No reference home exists
                Long upperReference = areaStatisticsMap.higherKey(reqArea);
                Long lowerReference = areaStatisticsMap.lowerKey(reqArea);
                if (upperReference == null) {
                    //reqArea is larger than all reference homes
                    upperReference = lowerReference;
                    lowerReference = areaStatisticsMap.lowerKey(upperReference);

                } else if (lowerReference == null) {
                    //reqArea is smaller than all reference homes
                    lowerReference = upperReference;
                    upperReference = areaStatisticsMap.higherKey(lowerReference);
                }
                SquareFootage upper = areaStatisticsMap.get(upperReference);
                SquareFootage lower = areaStatisticsMap.get(lowerReference);
                double y1 = lower.meanPrice;
                double y2 = upper.meanPrice;
                double x1 = lowerReference;
                double x2 = upperReference;
                double gradient = (y2 - y1) / (x2 - x1);
                output = y1 + (gradient * (reqArea - x1));

            }
        }
        output = Math.max(output, Math.pow(10d, 3d));
        output = Math.min(output, Math.pow(10d, 6d));
        return Math.round(output);
    }

    public static class SquareFootage {
        List<Long> prices;
        double sampleSize;
        double meanPrice;
        double stdDev;
        double var;

        public SquareFootage(List<Long> prices) {
            this.prices = prices;
            sampleSize = prices.size();
            double sum = 0d;
            for (Long l : prices) sum += l;
            meanPrice = sum / sampleSize;
            double SSE = 0d;
            for (Long l : prices) {
                SSE += Math.pow(l - meanPrice, 2);
            }
            var = (SSE / sampleSize);
            stdDev = Math.pow(var, 0.5d);
        }

        //Unsafe for objects with 1 price
        public SquareFootage outlierFilter() {
            List<Long> filteredPrices = new ArrayList<>();
            for (Long l : prices) {
                List<Long> excludedPrices = new ArrayList<>(prices);
                excludedPrices.remove(l);
                SquareFootage sf = new SquareFootage(excludedPrices);
                if (!sf.isOutlier(l)) {
                    filteredPrices.add(l);
                }
            }
            if (filteredPrices.size() == 0) return null;
            else return (new SquareFootage(filteredPrices));
        }

        public boolean isOutlier(Long area) {
            return (Math.abs(area - meanPrice) > (3 * stdDev));
        }
    }


    static Integer[] memory;

    public static int maximum_path(List<Integer> node_values) {
        // Write your code here
        //Assuming the input represents a valid triangular structure
        int n = node_values.size();
        if (n == 1) return node_values.get(0);
        memory = new Integer[n + 1];
        return maxPathSum(0, 0, 1, node_values, n);
    }

    public static int maxPathSum(int nodeIndex, int nodeOrder, int row, List<Integer> nodeValues, int nodeLength) {
        if (memory[nodeIndex + nodeOrder] != null) return memory[nodeIndex + nodeOrder];
        int output = nodeValues.get(nodeIndex + nodeOrder);
        if (nodeOrder + row >= nodeLength) return memory[nodeIndex + nodeOrder] = output;
        else {
            output += Math.max(maxPathSum(nodeIndex, nodeOrder + row, row + 1, nodeValues, nodeLength),
                    maxPathSum(nodeIndex + 1, nodeOrder + row, row + 1, nodeValues, nodeLength));
        }
        return memory[nodeIndex + nodeOrder] = output;
    }


    public static int dayOfYear(int day, int month, int year) {
        //Does not consider leap years
        int[] daysInMonth = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
        int dayOfYear = day;
        dayOfYear += year * 365;
        for (int i = 0; i < month - 1; i++) {
            dayOfYear += daysInMonth[i];
        }
        dayOfYear += day;
        return dayOfYear;
    }


    public static void calcMissing(List<String> readings) {
        // Write your code here
        List<Integer> timeIndex = new ArrayList<>();
        List<Double> samples = new ArrayList<>();
        int n = readings.size();
        for (int i = 0; i < n; i++) {
            String[] sArray = readings.get(i).split("\t");
            String sample = sArray[1];
            if (sample.startsWith("Miss")) {
                samples.add(i, null);
            } else {
                samples.add(i, Double.valueOf(sample));
            }
            sArray = sArray[0].split(" ");
            sArray = sArray[0].split("/");
            Integer time = dayOfYear(Integer.parseInt(sArray[0]), Integer.parseInt(sArray[1]), Integer.parseInt(sArray[2]));
            timeIndex.add(i, time);
        }
        if (samples.get(0) == null) {
            int index = 1;
            while (index < n && samples.get(index) == null) {
                index++;
            }
            Double d = samples.get(index);
            for (int i = 0; i < index; i++) {
                samples.set(i, d);
            }
        }
        if (samples.get(n - 1) == null) {
            int index = n - 2;
            while (index >= 0 && samples.get(index) == null) {
                index--;
            }
            Double d = samples.get(index);
            for (int i = index + 1; i < n; i++) {
                samples.set(i, d);
            }
        }
        for (int i = 0; i < n; i++) {
            if (samples.get(i) == null) {
                int left = i - 1;
                int right = i + 1;
                while (left >= 0 && samples.get(left) == null) {
                    left--;
                }
                while (right < n && samples.get(right) == null) {
                    right++;
                }
                if (left >= 0 && right < n) {
                    double x1 = timeIndex.get(left);
                    double y1 = samples.get(left);
                    double x2 = timeIndex.get(right);
                    double y2 = samples.get(right);
                    double interpol = linearInterpolate(x1, y1, x2, y2, timeIndex.get(i));
                    interpol = Math.max(interpol, 0);
                    interpol = Math.min(interpol, 399);
                    samples.set(i, interpol);
                    System.out.println(interpol);
                }
            }
        }
    }
    public static double linearInterpolate(double x1, double y1, double x2, double y2, double x) {
        return y1 + ((y2 - y1) / (x2 - x1)) * (x - x1);
    }


    public static void main(String[] args) {
        System.out.println("Hello world!");
    }
}