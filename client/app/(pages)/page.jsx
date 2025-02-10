"use client";
import { useEffect, useState } from 'react';
import pdfToText from 'react-pdftotext';

export default function Home() {
  const [pdfText, setPdfText] = useState("");
  const [loading, setLoading] = useState(false);

  function extractText(event) {
    const file = event.target.files[0];
    setLoading(true); // Start loading state

    pdfToText(file)
      .then(text => {
        setPdfText(text);
        setLoading(false); // Stop loading state
      })
      .catch(error => {
        console.error("Failed to extract text from pdf", error);
        setLoading(false); // Stop loading state even if there's an error
      });
  }

  // const sendPdfTextToBackend = async (pdfText) => {
  //   console.log("Sending PDF text to backend...");
  //   try {
  //     const resp = await fetch('http://localhost:8080/getResumeText', {
  //       method: 'POST',
  //       headers: {
  //         'Content-Type': 'application/json',
  //       },
  //       body: JSON.stringify({ text: pdfText }),
  //     });

  //     if (!resp.ok) {
  //       throw new Error("Network response was not ok");
  //     }

  //     const result = await resp.json();
  //     console.log("Response from backend:", result);
  //     sessionStorage.setItem('resumeText', JSON.stringify(result));

  //     return result;
  //   } catch (error) {
  //     console.error("Error sending PDF text:", error);
  //     return null;
  //   }
  // };

  useEffect(() => {
    if (pdfText) {
      console.log("PDF Text extracted:", pdfText);
      // sendPdfTextToBackend(pdfText);
    }
  }, [pdfText]);

  return (
    <>
      <div>
        <input type="file" accept="application/pdf" onChange={extractText} />
        {loading && <p>Loading...</p>} {/* Display loading state */}
      </div>
    </>
  );
}
