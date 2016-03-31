package tmp;

public class Transition {

    String event;
    String action;
    String stateid;

    public Transition(String event, String action, String stateid) {
        this.event = event;
        this.action = action;
        this.stateid = stateid;
    }

    public String getEvent() {
        return event;
    }

    public void setEvent(String event) {
        this.event = event;
    }

    public String getAction() {
        return action;
    }

    public void setAction(String action) {
        this.action = action;
    }

    public String getStateid() {
        return stateid;
    }

    public void setStateid(String stateid) {
        this.stateid = stateid;
    }
}
