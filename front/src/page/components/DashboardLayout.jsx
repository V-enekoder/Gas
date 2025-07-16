// Este componente de React define el layout principal del dashboard, incluyendo la barra lateral de navegación y un área para renderizar las rutas anidadas, además de un modal para completar el perfil del usuario.

import React from 'react';
import { Outlet } from 'react-router-dom';
import Sidebar from './Sidebar';
import '../Styles/dashboard.css';
import CompleteProfileModal from './CompleteProfileModal'
import { useAuth } from '../../context/AuthContext'; 

const DashboardLayout = () => {
    const { user } = useAuth(); 

    return (
        <>
          {user && !user.profileCompleted && <CompleteProfileModal />}

          <div className="dashboard-layout">
            <Sidebar />
            <main className="dashboard-main-content">
              <Outlet />
            </main>
          </div>
        </>
    );
};

export default DashboardLayout;