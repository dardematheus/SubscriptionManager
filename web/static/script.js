document.getElementById("registerUser").addEventListener('submit', function(event){
    event.preventDefault();
    
    const formData = new FormData(this);
    fetch("/register", {
        method: "POST",
        body: new URLSearchParams(formData)
    })
    .then(response => response.json())
    .then(data => {
        if (data.error){
            document.getElementById("err-msg").innerText = data.error;
        }
    })
})

document.getElementById("loginUser").addEventListener('submit', function(event){
    event.preventDefault();
    
    const formData = new FormData(this);
    fetch("/login", {
        method: "POST",
        body: new URLSearchParams(formData)
    })
    .then(response => response.json())
    .then(data => {
        if (data.error){
            document.getElementById("err-msg").innerText = data.error;
        }
    })
})