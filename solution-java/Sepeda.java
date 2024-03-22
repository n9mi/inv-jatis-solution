public class Sepeda extends Kendaraan {
    private String rantai;

    public Sepeda() {
        super("Swoosh");
        this.rantai = "Normal";
    }

    public void akselerasi() {
        super.akselerasi();
        this.rantai = "Perlu perbaikan";
    }

    public String getRantai() {
        return rantai;
    }
}
