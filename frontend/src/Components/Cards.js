import React, { useState } from 'react'
import PropTypes from 'prop-types'
import Select from 'react-select'
import Axios from 'axios';
const options=[
    //"red", "pink", "aqua", "blue", "brown", "olive", "green", "teal", "yellow", "fuchsia", "lime"
    {value: 1, label:'red'},
    {value: 2, label:'pink'},
    {value: 3, label:'aqua'},
    {value: 4, label:'blue'},
    {value: 5, label:'brown'},
    {value: 6, label:'olive'},
    {value: 7, label:'green'},
    {value: 8, label:'teal'},
    {value: 9, label:'yellow'},
    {value: 10, label:'fuchsia'},
    {value: 11, label:'lime'},
] 
const Cards = ({ name, Imagen, Color, Id }) => {
    const [valor,setValor]=useState(0)
    const [Colors, setColor] = useState('')
    const [Foto, setFoto] = useState('')
    const submitRegistro=()=>{
        if (valor===1){
            setColor('red');
        }else if (valor===2){
            setColor('pink');
        }else{
            setColor('teal')
        }
        


    }
    const [toggleState, setToggleState] = useState(1);
    const toggleTab = (index) => {
        setToggleState(index);
    }


    return (
        <React.Fragment>

            <div className="Titulo" >
                <h1> </h1>
            </div>
            <div className="containerCards">
                <div className="bloc-tabs">
                    <div className={toggleState === 1 ? "tabs active-tabs4" : "tabs"}
                        onClick={() => toggleTab(1)}>Deporte</div>
                    <div className={toggleState === 2 ? "tabs active-tabs4" : "tabs"}
                        onClick={() => toggleTab(2)}>Editar</div>
                    
                </div>
                <div className={toggleState === 1 ? "content active-content" : "content"}>

                    <div className="content active-content">
                        <div className="single-Deporte" style={{
                            backgroundColor: Color
                        }}>
                            <h2>{name}{Id}</h2>
                            <img src={`data:image/gif,base64,${Imagen}`} alt={name} />
                            

                        </div>

                    </div>
                </div>
                <div className={toggleState === 2 ? "content active-content" : "content"}>

                    <div className="content active-content">
                    <div className="form">
                <label>Color</label>
                <div clasName="SelectColor"><Select  options={options} onChange={(e) =>
                    setValor(e.target.value)
                } />
                </div>
                
                <label>Foto</label>
                <input type="file" name="Foto" onChange={(e) =>
                    setFoto(e.target.files[0].name)
                } />
                <button onClick={submitRegistro}>Editar</button>
            </div>
                    </div>
                </div>
               
            </div>


        </React.Fragment>
    );
}

Cards.propTypes = {
    name: PropTypes.string
}
export default Cards;