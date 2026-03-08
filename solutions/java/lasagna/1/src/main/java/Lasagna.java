public class Lasagna {
    public int expectedMinutesInOven() {
        return 40;
    }

    public int remainingMinutesInOven(int actualMinutes) {
        int expectedMinutesInOven = expectedMinutesInOven();
        return expectedMinutesInOven - actualMinutes;
    }

    public int preparationTimeInMinutes(int numberOfLayers ) {
        return numberOfLayers * 2;
    }

    public int totalTimeInMinutes(int numberOfLayers, int numberOfMinutes) {
        int preparationTime = preparationTimeInMinutes(numberOfLayers);
        return preparationTime + numberOfMinutes;
    }
}