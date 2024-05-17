// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyBjIo10JsSYzw3D-3VId8DkN5TgHlKqGtI",
  authDomain: "test-submsg.firebaseapp.com",
  projectId: "test-submsg",
  storageBucket: "test-submsg.appspot.com",
  messagingSenderId: "777680085082",
  appId: "1:777680085082:web:3938cbd9935c52a1283863",
  measurementId: "G-777SNVC3LP"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);