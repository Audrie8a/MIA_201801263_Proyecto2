import React from 'react'
import '../assets/css/Tier.css'
import Card from './Cards'
import Axios from 'axios'
import '../assets/css/Deporte.css'
class Deporte extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            Deportes: [],
        }
    }
    componentDidMount() {
        Axios.get('http://localhost:4000/Deportes').then((response) => {

            this.setState({ Deportes: response.data.Sports })

        })
    }
    handleChange(checked) {
        this.setState({ checked });
      }
    render() {
        const { Deportes } = this.state
        return (
            <React.Fragment>
            <div className="Titulos">
                    <h1>Lista de Deportes</h1>
                    <br/>
            </div>
            <section className="Deportes-container">
                
                {
                    Deportes.map((Deporte) => <Card
                        name={Deporte.Nombre}
                        Imagen={Deporte.Imagen}
                        Color={Deporte.Color}
                        Id= {Deporte.IdDeporte}
                    />)
                }
            </section>
            
           
   
           
            </React.Fragment>
        )
    }
}

export default Deporte;