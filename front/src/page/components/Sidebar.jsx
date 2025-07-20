// Este componente de React renderiza la barra lateral de navegación del dashboard, proporcionando enlaces a las diferentes secciones de la aplicación y gestionando la funcionalidad de cierre de sesión del usuario.

import React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import { FaTachometerAlt, FaBoxOpen, FaTruck, FaFileAlt, FaSignOutAlt, FaUser } from 'react-icons/fa';
import { useAuth } from '../../context/AuthContext';
import '../Styles/dashboard.css';

const Sidebar = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = (e) => {
    e.preventDefault();
    logout();
    navigate('/login');
  };

  return (
    <aside className="sidebar">
      <div className="sidebar-header">
        <h3>Sistema de Gas</h3>
      </div>
      <nav className="sidebar-nav">
        <ul>
          <li><NavLink to="/dashboard"><FaTachometerAlt /> Dashboard</NavLink></li>
          <li><NavLink to="/orders"><FaBoxOpen /> Órdenes</NavLink></li>
          <li><NavLink to="/deliveries"><FaTruck /> Entregas</NavLink></li>
          <li><NavLink to="/reports"><FaFileAlt /> Reportes</NavLink></li>
          <li><NavLink to="/my-data"><FaUser /> Mis Datos</NavLink></li>
        </ul>
      </nav>
      <div className="sidebar-footer">
        {user && (
          <>
            <div className="user-info">
              <span>{user.name}</span>
              <a href="/logout" onClick={handleLogout} className="logout-link">
                <FaSignOutAlt /> Cerrar Sesión
              </a>
            </div>
          </>
        )}
      </div>
    </aside>
  );
};

export default Sidebar;