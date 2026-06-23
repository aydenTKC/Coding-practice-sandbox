interface BusinessPartner {
    name: string;
}

interface ContactDetail {
    email: string;
    phone: string;
}

type BusinessContact = BusinessPartner & ContactDetail

/*
const contact: BusinessContact = {
    name: "Ayden",
    email: "ayden00lee@gmail.com",
    phone: "6514700753"
}
 */
