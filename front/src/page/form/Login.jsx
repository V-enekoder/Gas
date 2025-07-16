// Este componente de React gestiona el inicio de sesión de usuarios, autenticándolos con la API, actualizando el estado global de la aplicación y redirigiéndolos al dashboard.

import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FaEnvelope, FaLock, FaEye, FaEyeSlash   } from 'react-icons/fa';
import { useAuth } from '../../context/AuthContext';
import { authService } from '../../services/ApiServices';
import '../Styles/form.css';


const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false); 
  
  const navigate = useNavigate();
  const { login: loginContext } = useAuth();

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      const response = await authService.login({ email, password });
      const userDataFromAPI = response.data;
      
      console.log("Respuesta de la API en Login:", userDataFromAPI);

      if (!userDataFromAPI || !userDataFromAPI.id || !userDataFromAPI.role) {
        throw new Error("La respuesta de la API es inválida o no contiene los datos necesarios.");
      }

      const isProfileCompleted = userDataFromAPI.role === 'user' || (userDataFromAPI.role_data != null);
      
      if (userDataFromAPI.token) {
        localStorage.setItem('authToken', userDataFromAPI.token);
      }
      
      const userForContext = {
        ...userDataFromAPI,
        isAuthenticated: true,
        profileCompleted: isProfileCompleted,
      };

      loginContext(userForContext);

      navigate('/dashboard');

    } catch (err) {
      const errorMessage = err.response?.data?.error || 'Credenciales inválidas o error en el servidor.';
      setError(errorMessage);
      console.error("Error al iniciar sesión:", err.response || err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const togglePasswordVisibility = () => {
    setShowPassword(prevShowPassword => !prevShowPassword);
  };

  return (
    <div className="auth-page-container">
      <div className="login-card">
        <h1>Sistema de Gas</h1>
        <p>Bienvenido, por favor inicia sesión.</p>

        <form onSubmit={handleSubmit}>
           <div className="form-group">
            <input type="email" id="email" className="form-input" placeholder=" " value={email} onChange={(e) => setEmail(e.target.value)} required disabled={isLoading} />
            <label htmlFor="email" className="form-label">Correo Electrónico</label>
            <FaEnvelope className="input-icon" />
          </div>

          <div className="form-group">
            <input 
              type={showPassword ? "text" : "password"} 
              id="password" 
              className="form-input" 
              placeholder=" " 
              value={password} 
              onChange={(e) => setPassword(e.target.value)} 
              required 
              disabled={isLoading} 
            />
            <label htmlFor="password" className="form-label">Contraseña</label>
            <FaLock className="input-icon" />

            <span className="password-toggle-icon" onClick={togglePasswordVisibility}>
              {showPassword ? <FaEyeSlash /> : <FaEye />}
            </span>
          </div>

          {error && <p className="error-message">{error}</p>}
          <button type="submit" className="btn-login" disabled={isLoading}>
            {isLoading ? 'Iniciando...' : 'Iniciar Sesión'}
          </button>
        </form>

        <div className="auth-switch-link">
          ¿No tienes una cuenta?{' '}
          <Link to="/register">Regístrate aquí</Link>
        </div>
      </div>
    </div>
  );
};

export default Login;