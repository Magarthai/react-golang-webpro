import React, {useState} from 'react'
import "/src/CSS/Menubra.css"
import { Link } from 'react-router-dom'
import { FiMenu, FiX} from "react-icons/fi";
import { FaUserCircle } from 'react-icons/fa';


function menubra() {

    const [click, setclick] = useState(false);
    const handleClick = () => setclick(!click);
    const closeMenu = () => setclick(false);

  return (
    <header className="header">
        <section className="container">
            <div className="menu-con">
                <div className="logo">
                    <a href="#">
                        <img src="./src/Image/logo2.jpg" alt="TC GarageCLTB" />
                    </a>
                </div>
                <ul className= {click ? "menu active" : "menu"}>
                    <li className="menulink" onClick={closeMenu}>
                        <Link to = "/">หน้าหลัก</Link>
                    </li>
                    <li className="menulink" onClick={closeMenu}>
                        <Link to = "/gallery">แกลเลอรี่</Link>
                    </li>
                    <li className="menulink" onClick={closeMenu}>
                        <Link to = "/profile">โปรไฟล์</Link>
                    </li>
                    <li className="menulink" onClick={closeMenu}>
                        <Link to = "/news">ข่าวสาร</Link>
                    </li>
                    <li className="menulink" onClick={closeMenu}>
                        <Link to = "/member">สมาชิก</Link>
                    </li>
                    <li className="menulink" onClick={closeMenu}>
                        <Link to = "/about">เกี่ยวกับ</Link>
                    </li>
                    <li className="menulink" id="User" onClick={closeMenu}>
                        <Link to = "/signup"><FaUserCircle /></Link>
                    </li>
                </ul>

                <div className="menumobile" onClick={handleClick}>
                    {click ? (
                        <FiX />
                    ) : (
                        <FiMenu />
                    )}
                </div>
            </div>
        </section>
    </header>
  )
}

export default menubra