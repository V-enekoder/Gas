// Este componente de React permite a los usuarios registrarse en la aplicación, solicitando información básica, seleccionando un municipio y un tipo de rol, y gestionando el proceso de registro con la API.

import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { authService, municipalityService } from '../../services/ApiServices';
import { useAuth } from '../../context/AuthContext';
import '../Styles/form.css';
import { FaUser, FaStore, FaUsers, FaWheelchair, FaUserCircle, FaEnvelope, FaLock, FaMapMarkerAlt, FaEye, FaEyeSlash, FaHome } from 'react-icons/fa';


const roleOptions = [
  { value: 'user', label: 'Estándar', icon: FaUser },
  { value: 'commerce', label: 'Comercio', icon: FaStore },
  { value: 'council', label: 'Comunal', icon: FaUsers },
  { value: 'disabled', label: 'Discapacidad', icon: FaWheelchair },
];

const Register = () => {
  const [fullName, setFullName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('user');
  const [showPassword, setShowPassword] = useState(false);


  const [municipalities, setMunicipalities] = useState([]);
  const [selectedMunicipality, setSelectedMunicipality] = useState('');

  const [error, setError] = useState(null);
  const [emailError, setEmailError] = useState('');
  const [passwordError, setPasswordError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();
  const { login } = useAuth();

  useEffect(() => {
    const fetchMunicipalities = async () => {
      try {
        const response = await municipalityService.getAll();
        setMunicipalities(response.data);
        if (response.data.length > 0) {
          setSelectedMunicipality(response.data[0].id);
        }
      } catch (err) {
        console.error("Error al obtener los municipios:", err);
        setError("No se pudieron cargar los municipios. Inténtalo de nuevo más tarde.");
      }
    };

    fetchMunicipalities();
  }, []);

  const validateEmailFormat = (emailValue) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(emailValue)) {
      setEmailError('Formato de correo inválido.');
      return false;
    }
    setEmailError('');
    return true;
  };

  const validatePasswordLength = (passwordValue) => {
    if (passwordValue.length < 8) {
      setPasswordError('La contraseña debe tener al menos 8 caracteres.');
      return false;
    }
    setPasswordError('');
    return true;
  };

  const handleEmailChange = (e) => {
    const value = e.target.value;
    setEmail(value);
    validateEmailFormat(value);
  };

  const handlePasswordChange = (e) => {
    const value = e.target.value;
    setPassword(value);
    validatePasswordLength(value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const isEmailValid = validateEmailFormat(email);
    const isPasswordValid = validatePasswordLength(password);

    if (!isEmailValid || !isPasswordValid) {
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      const registrationData = {
        name: fullName,
        email,
        password,
        municipality_id: parseInt(selectedMunicipality, 10),
      };

      const response = await authService.register(registrationData);
      const userDataFromAPI = response.data;

      if (userDataFromAPI.token) {
        localStorage.setItem('authToken', userDataFromAPI.token);
      }

      const userForContext = {
        ...userDataFromAPI,
        isAuthenticated: true,
        role: role,
        profileCompleted: role === 'user',
      };

      login(userForContext);
      navigate('/dashboard');

    } catch (err) {
      const errorMessage = err.response?.data?.error || 'No se pudo completar el registro. Inténtalo de nuevo.';
      setError(errorMessage);
      console.error("Error en el registro:", err.response || err);
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
        <h1>Crear una Cuenta</h1>
        <p>Únete a nuestro sistema de gestión de gas.</p>

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <input type="text" id="fullName" className="form-input" placeholder=" " value={fullName} onChange={(e) => setFullName(e.target.value)} required />
            <label htmlFor="fullName" className="form-label">Nombre y Apellido</label>
            <FaUserCircle className="input-icon" />
          </div>

          <div className="form-group">
            <input type="email" id="email" className="form-input" placeholder=" " value={email} onChange={handleEmailChange} required />
            <label htmlFor="email" className="form-label">Correo Electrónico</label>
            <FaEnvelope className="input-icon" />
            {emailError && <p className="input-error-message">{emailError}</p>}
          </div>

          <div className="form-group">
            <input type={showPassword ? "text" : "password"} id="password" className="form-input" placeholder=" " value={password}
              onChange={handlePasswordChange} required disabled={isLoading} />
            <label htmlFor="password" className="form-label">Contraseña</label>
            <FaLock className="input-icon" />
            <span className="password-toggle-icon" onClick={togglePasswordVisibility}>
              {showPassword ? <FaEyeSlash /> : <FaEye />}
            </span>

            {passwordError && <p className="input-error-message">{passwordError}</p>}
          </div>

          <div className="form-group">
            <FaMapMarkerAlt className="select-icon" />
            <select
              id="municipality"
              className="form-select"
              value={selectedMunicipality}
              onChange={(e) => setSelectedMunicipality(e.target.value)}
              required
              disabled={municipalities.length === 0}
            >
              {municipalities.length === 0 ? (
                <option>Cargando municipios...</option>
              ) : (
                municipalities.map(muni => (
                  <option key={muni.id} value={muni.id}>
                    {muni.name}
                  </option>
                ))
              )}
            </select>
          </div>

          <div className="form-group">
            <label className="static-label">Tipo de Usuario</label>
            <div className="role-selector">
              {roleOptions.map((option) => (
                <div key={option.value}>
                  <input type="radio" id={`role-${option.value}`} name="role" value={option.value} checked={role === option.value} onChange={() => setRole(option.value)} className="role-input-hidden" />
                  <label htmlFor={`role-${option.value}`} className="role-option-card">
                    <option.icon className="role-icon" />
                    <span>{option.label}</span>
                  </label>
                </div>
              ))}
            </div>
          </div>

          {error && <p className="error-message">{error}</p>}

          <button type="submit" className="btn-login" disabled={isLoading}>
            {isLoading ? 'Registrando...' : 'Registrarse'}
          </button>
        </form>

        <div className="auth-switch-link">
          ¿Ya tienes una cuenta?{' '}
          <Link to="/login">Inicia sesión aquí</Link>
        </div>
      </div>
      <button onClick={() => navigate('/')} className="floating-home-btn">
        <FaHome />
      </button>

    </div>
  );
};

export default Register;