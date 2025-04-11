document.addEventListener('DOMContentLoaded', function() {
    const personForm = document.getElementById('personForm');
    const totalHombres = document.getElementById('totalHombres');
    const totalMujeres = document.getElementById('totalMujeres');
    const personasList = document.getElementById('personas-list');

    // Recuperar contadores del localStorage si existen
    let hombresCount = parseInt(localStorage.getItem('totalHombres')) || 0;
    let mujeresCount = parseInt(localStorage.getItem('totalMujeres')) || 0;

    // Actualizar los contadores al cargar la página
    totalHombres.textContent = hombresCount;
    totalMujeres.textContent = mujeresCount;

    // Recuperar la lista de personas
    const personasGuardadas = JSON.parse(localStorage.getItem('personas')) || [];

    // Mostrar las personas guardadas
    function renderPersonas() {
        personasList.innerHTML = '';
        personasGuardadas.forEach(persona => {
            const li = document.createElement('li');
            li.textContent = `${persona.nombre} - ${persona.edad} años - ${persona.sexo}`;
            personasList.appendChild(li);
        });
    }

    renderPersonas();

    // Evento de envío del formulario
    personForm.addEventListener('submit', function(e) {
        e.preventDefault();

        const nombre = document.getElementById('nombre').value;
        const edad = parseInt(document.getElementById('edad').value);
        const sexo = document.querySelector('input[name="sexo"]:checked').value;

        // Crear el objeto persona
        const persona = {
            nombre: nombre,
            edad: edad,
            sexo: sexo
        };

        // Guardar persona en la lista
        personasGuardadas.push(persona);
        localStorage.setItem('personas', JSON.stringify(personasGuardadas));

        // Actualizar los contadores
        if (sexo === 'hombre') {
            hombresCount++;
            localStorage.setItem('totalHombres', hombresCount);
            totalHombres.textContent = hombresCount;
        } else if (sexo === 'mujer') {
            mujeresCount++;
            localStorage.setItem('totalMujeres', mujeresCount);
            totalMujeres.textContent = mujeresCount;
        }

        // Limpiar el formulario
        personForm.reset();

        // Volver a renderizar la lista de personas
        renderPersonas();
    });
});
