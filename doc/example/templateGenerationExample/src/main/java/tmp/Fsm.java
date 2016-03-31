package tmp;

import java.util.ArrayList;

public class Fsm {

    ArrayList<Statedecl> statedecl;

    public Fsm(ArrayList<Statedecl> statedecl) {
        this.statedecl = statedecl;
    }

    public ArrayList<Statedecl> getStatedecl() {
        return statedecl;
    }

    public void setStatedecl(ArrayList<Statedecl> statedecl) {
        this.statedecl = statedecl;
    }
}
