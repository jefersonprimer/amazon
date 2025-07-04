import { Metadata } from "next";
import Header from "./components/layout/header";
import HeaderMain from "./components/layout/HeaderMain";


const Home = () => {
    return (
        <>
            <Header />
            <HeaderMain/>
            <main style={{ padding: "2rem", textAlign: "center" }}>
                <h1>Bem-vindo à Frederico Westphalen</h1>
                <p>Explore nossos produtos, categorias e muito mais!</p>
            </main>
        </>
    );
}

export default Home;