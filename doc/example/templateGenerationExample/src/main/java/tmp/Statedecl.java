package tmp;

import java.util.ArrayList;

public class Statedecl {

    boolean initial;
    String stateid;
    ArrayList<Transition> transition;


    public Statedecl(boolean initial, String stateid, ArrayList<Transition> transition) {
        this.initial = initial;
        this.stateid = stateid;
        this.transition = transition;
    }

    public boolean isInitial() {
        return initial;
    }

    public void setInitial(boolean initial) {
        this.initial = initial;
    }

    public String getStateid() {
        return stateid;
    }

    public void setStateid(String stateid) {
        this.stateid = stateid;
    }

    public ArrayList<Transition> getTransition() {
        return transition;
    }

    public void setTransition(ArrayList<Transition> transition) {
        this.transition = transition;
    }
}
