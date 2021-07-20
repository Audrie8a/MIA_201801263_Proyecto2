import React from 'react'
import '../assets/css/Tier.css'
import Card from './Cards'
import Axios from 'axios'
import '../assets/css/Deporte.css'

class Eventos extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            Deportes: [],
        }
    }
    componentDidMount() {
        Axios.get('http://localhost:4000/Eventos').then((response) => {

            this.setState({ Deportes: response.data.Eventoss })

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
                        IdEvento={Deporte.Eventoss}
                        IdJornada={Deporte.IdJornada}
                        Nombre={Deporte.Nombre}
                        Fecha= {Deporte.Fecha}
                        NombreLocal= {Deporte.NombreLocal}
                        NombreVisitante= {Deporte.NombreVisitante}
                        
                    />)
                }
            </section>
            
           
   
           
            </React.Fragment>
        )
    }
}

export default Eventos;