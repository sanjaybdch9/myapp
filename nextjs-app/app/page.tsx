import Image from "next/image";
import styles from "./page.module.css";

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <p>
          Hello, World! This is my first nextjs application.&nbsp;
        </p>
      
      </div>
    </main>
  );
}
