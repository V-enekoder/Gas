import axios from 'axios';

// Configura el cliente de API con la URL base.
const apiClient = axios.create({
    baseURL: 'http://localhost:8080', 
    headers: {
        'Content-Type': 'application/json',
    }
});

// Intercepta peticiones para añadir el token de autenticación.
apiClient.interceptors.request.use(
    config => {
        const token = localStorage.getItem('authToken');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => Promise.reject(error)
);

export const authService = {
    // Inicia sesión de usuario.
    login: (credentials) => apiClient.post('/users/login', credentials),
    // Registra un nuevo usuario.
    register: (userData) => apiClient.post('/users/register', userData),
};

export const profileService = {
    // Completa el perfil de un comercio.
    completeCommerceProfile: (data) => apiClient.post('/commerces/', data),
    // Completa el perfil de un concejo comunal.
    completeCouncilProfile: (data) => apiClient.post('/councils/', data),
    // Completa el perfil de una persona con discapacidad.
    completeDisabledProfile: (data) => apiClient.post('/disabled/', data),
};

export const municipalityService = {
    // Obtiene todos los municipios.
    getAll: () => apiClient.get('/municipalities/'),
};

export const orderService = {
    // Crea una nueva orden.
    create: (orderData) => apiClient.post('/orders/', orderData),
    // Obtiene una orden por su ID.
    getById: (orderId) => apiClient.get(`/orders/${orderId}`),
    // Obtiene órdenes por ID de usuario.
    getByUserId: (userId) => apiClient.get(`/orders/${userId}/user`),

};

export const cylinderService = {
    // Obtiene todos los tipos de cilindros.
    getAllTypes: () => apiClient.get('/type-cylinders/'), 
};

export const paymentService = {
    // Crea un nuevo registro de pago.
    create: (paymentData) => apiClient.post('/payments/', paymentData),
};

export const deliveryService = {
    // Obtiene entregas por ID de usuario.
    getByUserId: (userId) => apiClient.get(`/deliveries/${userId}/user`),
    // Obtiene una entrega por su ID.
    getById: (deliveryId) => apiClient.get(`/deliveries/${deliveryId}`),
    // Crea una nueva entrega.
    create: (data) => apiClient.post('/deliveries/', data),
};

export const reportService = {
    // Obtiene reportes por ID de usuario.
    getByUserId: (userId) => apiClient.get(`/reports/${userId}/user`),
    // Obtiene todos los tipos de reporte.
    getReportTypes: () => apiClient.get('/report-types/'),
    // Obtiene todos los estados de reporte.
    getReportStates: () => apiClient.get('/report-states/'),
    // Crea un nuevo reporte.
    create: (reportData) => apiClient.post('/reports/', reportData), 
};