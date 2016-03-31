import org.stringtemplate.v4.ST;
import org.stringtemplate.v4.STGroup;
import org.stringtemplate.v4.STGroupFile;
import  tmp.*;

import java.util.ArrayList;
import java.util.Arrays;

public class Main {

    public static void main(String[] args){
        Fsm fsm = new Fsm(new ArrayList());
        fsm.getStatedecl().add(
                new Statedecl(
                        true,
                        "state_0",
                        new ArrayList<Transition>(
                                Arrays.asList(
                                        new Transition(
                                                "event0_0",
                                                "action0_1",
                                                "state_1"
                                        ),
                                        new Transition(
                                                "event0_1",
                                                null,
                                                null
                                        ),
                                        new Transition(
                                                "event0_2",
                                                "action0_11",
                                                "state_1"
                                        )
                                )
                        )
                )
        );

        fsm.getStatedecl().add(
                new Statedecl(
                        false,
                        "state_1",
                        new ArrayList<Transition>(
                                Arrays.asList(
                                        new Transition(
                                                "event1_1",
                                                "action1_2",
                                                "state_2"
                                        ),
                                        new Transition(
                                                "event1_2",
                                                 null,
                                                "state_3"
                                        )
                                )
                        )
                )
        );

        fsm.getStatedecl().add(
                new Statedecl(
                        false,
                        "state_2",
                        new ArrayList<Transition>(
                                Arrays.asList(
                                        new Transition(
                                                "event2_1",
                                                "action2_1",
                                                "state_2"
                                        ),
                                        new Transition(
                                                "event2_2",
                                                null,
                                                "state_3"
                                        )
                                )
                        )
                )
        );

        STGroup stGroup = new STGroupFile("src/main/java/tmp/temp.stg");
        ST st = stGroup.getInstanceOf("fsm");
        st.add("fsm", fsm);
        System.out.print(st.render());
    }
}
