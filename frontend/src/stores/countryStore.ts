// import { defineStore } from 'pinia'

// interface ICountryState {
//   countries: Array<ICountry>;
// }

// export const useCountryStore = defineStore("countryStore", {
//   state: (): ICountryState => {
//     return {
//       countries: [],
//     };
//   },
//   getters: {
//     getCountriesWithFlags(): IMinimalCountryDetails {
//       return this.countries.map((item) => {
//         const name = item.name.common;
//         const icon = item.flags;
//         return { icon: icon.svg || icon.png, name };
//       });
//     },
//     hasData(): boolean {
//       return this.countries.length > 0 ? true : false;
//     },
//   },
//   actions: {
//     loadAllCountries() {
//       const { data, error } = useFetch(
//         useRuntimeConfig().public.COUNTRY_URI_API
//       );

//       if (data) {
//         // console.log("Hole Daten", data);
//         this.countries = data.value;
//       }
//       if (error) {
//         this.countries = [];
//       }
//     },
//     clearUser() {
//       this.$reset();
//     },
//   },
// });
