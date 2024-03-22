public class Main {
    public static void main(String[] args) {
        Sepeda sepeda = new Sepeda();
        sepeda.akselerasi();
        System.out.println(sepeda.getRantai());

        Mobil mobil = new Mobil();
        mobil.akselerasi();
        System.out.println(mobil.getBensin());
    }
}
