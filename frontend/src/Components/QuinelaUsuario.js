import React from 'react'
import '../assets/css/Tier.css'
import FilaQuinelaUsuario from './FilaQuinelaUsuario'
import Axios from 'axios'
import '../assets/css/Deporte.css'


class QuinelaUsuario extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            Quinelas: [],
        }
    }
    componentDidMount() {       
        Axios.get('http://localhost:4000/Quinelas').then((response) => {

                this.setState({ Quinelas: response.data.Quinelas })

            })
    }
    handleChange(checked) {
        this.setState({ checked });
    }
    render() {
        const { Quinelas } = this.state
        return (
            <React.Fragment>
                <div className="Titulos">
                    <h1>Quinelas</h1>
                    <br />
                    
                </div>
                <table className="table table-bordered">
                    <thead>
                        <tr>
                            <th>Username</th>
                            <th>Id Quinela</th>
                            <th>Puntaje</th>
                            <th>Nombre</th>
                            <th>Equipo Local</th>
                            <th>Resultado Local</th>
                            <th>Prediccion Local</th>
                            <th>Equipo Visitante</th>
                            <th>Resultado Visitante</th>
                            <th>Prediccion Visitante</th>
                            <th>Fecha</th>

                        </tr>
                    </thead>
                    <tbody>
                            {
                                Quinelas.map((Quinela) => <FilaQuinelaUsuario
                                    Username={Quinela.Username}
                                    IdQuinela={Quinela.IdQuinela}
                                    Puntaje={Quinela.Puntaje}
                                    Nombre={Quinela.Nombre}
                                    Local={Quinela.Local}
                                    Visitante={Quinela.Visitante}
                                    NombreVisitante={Quinela.NombreVisitante}
                                    NombreLocal={Quinela.NombreLocal}
                                    ResultadoVisitante={Quinela.ResultadoVisitante}
                                    ResultadoLocal={Quinela.ResultadoLocal}
                                    Fecha={Quinela.Fecha}

                                />)
                            }
                    </tbody>
                </table>





            </React.Fragment>
        )
    }
}

export default QuinelaUsuario;