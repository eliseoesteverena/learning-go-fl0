document.addEventListener("DOMContentLoaded", () => {
        const $btnEnviar = document.querySelector("#preview_btn"),
            $inputNameRt = document.querySelector("#name_rt"),
            $inputAddressRt = document.querySelector("#address_rt"),
            $inputCpRt = document.querySelector("#cp_rt"),
            $inputCityRt = document.querySelector("#city_rt"),
            $inputStateRt = document.querySelector("#state_rt"),
            $inputNameDt = document.querySelector("#name_dt"),
            $inputAddressDt = document.querySelector("#address_dt"),
            $inputCpDt = document.querySelector("#cp_dt"),
            $inputCityDt = document.querySelector("#city_dt"),
            $inputStateDt = document.querySelector("#state_dt");

    
        $btnEnviar.onclick = async () => {
            const nameRt = $inputNameRt.value;
            const addressRt = $inputAddressRt.value;
            const cpRt = $inputCpRt.value;
            const cityRt = $inputCityRt.value;
            const stateRt = $inputStateRt.value;
            const nameDt = $inputNameDt.value;
            const addressDt = $inputAddressDt.value;
            const cpDt = $inputCpDt.value;
            const cityDt = $inputCityDt.value;
            const stateDt = $inputStateDt.value;
            
            var datos = {
                remitente: {
                    name: nameRt,
                    adress: addressRt,
                    cp: cpRt,
                    city:  cityRt,
                    state: stateRt
                },
                destinatario: {
                    name: nameDt,
                    adress: addressDt,
                    cp: cpDt,
                    city:  cityDt,
                    state: stateDt
                }
            }
            
            const URL_SERVIDOR = "./static"; // Servidor de Go
            const datosJson =  JSON.stringify(datos);
            console.log(datos);
            try {
                const response = await fetch(URL_SERVIDOR, {
                    method: "POST",
                    body: datosJson,
                });
                console.log(response);
                const respuesta = await response.text();
                console.log("El servidor dijo: " + respuesta)
            } catch (e) {
                console.log("Error en el servidor: " + e.message);
            } 
        };

    });
