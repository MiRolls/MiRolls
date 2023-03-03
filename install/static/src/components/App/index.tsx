import {Route, Routes} from "react-router-dom";
import FirstStep from "../pages/FirstStep";
import SecondStep from "../pages/SecondStep";
import ThirdStep from "../pages/ThirdStep";
import HelloPage from "../pages/HelloPage";
import Done from "../pages/Done";


export default function (){
    return(
        <div style={{height:"100vh",width:"100%"}}>
            <Routes>
                <Route path="/" element={<HelloPage/>}/>
                <Route path="/step-one" element={<FirstStep/>}></Route>
                <Route path="/step-two" element={<SecondStep/>}></Route>
                <Route path="/step-three" element={<ThirdStep/>}></Route>
                <Route path="/done" element={<Done />}></Route>
            </Routes>
            {/*<Steps step={3} CompletedStep={0}></Steps>*/}
        </div>
    )
}
