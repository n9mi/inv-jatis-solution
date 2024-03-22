public class Mobil extends Kendaraan {
    private String bensin;

    public Mobil() {
        super("Vroom");
        this.bensin = "Penuh";
    }

    public void akselerasi() {
        super.akselerasi();
        this.bensin = "Kosong";
    }

    public String getBensin() {
        return bensin;
    }
}
