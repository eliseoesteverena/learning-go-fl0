    document.addEventListener("DOMContentLoaded", () => {
        const $inputName = document.querySelector("#nombre_rt"),
            $inputDire = document.querySelector("#domicilio_rt"),
            $btnEnviar = document.querySelector("#preview_btn");

        $btnEnviar.onclick = async () => {
            const nombre = $inputName.value;
            const direccion = $inputDire.value;
            var datos = {
                nombre: nombre,
                direccion: direccion
            }
            const URL_SERVIDOR = "./static"; // Servidor de Go
            const datosJson =  JSON.stringify(datos);
            console.log(datos);
            try {
                const response = await fetch(URL_SERVIDOR, {
                    method: "POST",
                    body: datosJson,
                });
                const respuesta = await response.text();
                console.log("El servidor dijo: " + respuesta)
            } catch (e) {
                console.log("Error en el servidor: " + e.message);
            }
        };

    });
