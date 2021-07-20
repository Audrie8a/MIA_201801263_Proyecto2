import React from 'react'
import '../assets/css/Tier.css'
import Axios from 'axios'
import '../assets/css/Deporte.css'
import FilaTemp from './FilaTemp'


class Temporadas extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            Quinelas: [],
        }
    }
    componentDidMount() {       
        Axios.get('http://localhost:4000/Temporadas').then((response) => {

                this.setState({ Quinelas: response.data.Datos })

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
                    <h1>Temporadas</h1>
                    <br />
                    
                </div>
                <table className="table table-bordered">
                    <thead>
                        <tr>
                            <th>Usuario</th>
                            <th>Temporada</th>
                            <th>Total</th>

                        </tr>
                    </thead>
                    <tbody>
                            {
                                Quinelas.map((Quinela) => <FilaTemp
                                    Username={Quinela.Username}
                                    Temp={Quinela.Nombre}
                                    Total={Quinela.Total}
                                    

                                />)
                            }
                    </tbody>
                </table>





            </React.Fragment>
        )
    }
}

export default Temporadas;