package main

import (
  "log"
  "earthcube.org/Project418/crawler/indexer"
  //"earthcube.org/Project418/crawler/framing"
)

func main() {
  //framing.SpatialFrame(iedaJsonld)
  //framing.SpatialFrame(bcodmoJsonld)
  msg := indexer.SpatialIndexer(bcodmoUrl, bcodmoJsonld)
  log.Println(msg)
  msg = indexer.SpatialIndexer(iedaUrl, iedaJsonld)
  log.Println(msg)
}

const bcodmoUrl string = "https://www.bco-dmo.org/dataset/472032"
const bcodmoJsonld string = `{"@context":{"@vocab":"http:\/\/schema.org\/","gdx":"https:\/\/geodex.org\/voc\/","datacite":"http:\/\/purl.org\/spar\/datacite\/","earthcollab":"https:\/\/library.ucar.edu\/earthcollab\/schema#","geolink":"http:\/\/schema.geolink.org\/1.0\/base\/main#","geolink-vocab":"http:\/\/schema.geolink.org\/1.0\/voc\/local#","vivo":"http:\/\/vivoweb.org\/ontology\/core#"},"@id":"https:\/\/www.bco-dmo.org\/dataset\/472032","identifier":["http:\/\/lod.bco-dmo.org\/id\/dataset\/472032",{"@type":"PropertyValue","additionalType":["http:\/\/schema.geolink.org\/1.0\/base\/main#Identifier","http:\/\/purl.org\/spar\/datacite\/Identifier"],"@id":"https:\/\/doi.org\/10.1575\/1912\/bco-dmo.665253","propertyID":"http:\/\/purl.org\/spar\/datacite\/doi","value":"10.1575\/1912\/bco-dmo.665253","url":"https:\/\/doi.org\/10.1575\/1912\/bco-dmo.665253"}],"url":"https:\/\/www.bco-dmo.org\/dataset\/472032","@type":"Dataset","name":"Removal of organic carbon by natural bacterioplankton communities as a function of pCO2 from laboratory experiments between 2012 and 2016","additionalType":["http:\/\/schema.geolink.org\/1.0\/base\/main#Dataset","http:\/\/vivoweb.org\/ontology\/core#Dataset"],"alternateName":"Data Set 3A: Utilization of dissolved organic carbon by a natural bacterial community as a function of pCO2","description":null,"datePublished":"2016-12-05","keywords":"ocean acidification, OA, Dissolved Organic Carbon, DOC, bacterioplankton respiration, pCO2, carbon dioxide, elevated pCO2, oceans","creator":[{"@type":"Role","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Participant","roleName":"Principal Investigator","creator":{"@type":"Person","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Person","@id":"https:\/\/www.bco-dmo.org\/person\/51317","name":"Dr Uta Passow","url":"https:\/\/www.bco-dmo.org\/person\/51317"}},{"@type":"Role","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Participant","roleName":"Co-Principal Investigator","creator":{"@type":"Person","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Person","@id":"https:\/\/www.bco-dmo.org\/person\/50575","name":"Dr Craig Carlson","url":"https:\/\/www.bco-dmo.org\/person\/50575"}},{"@type":"Role","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Participant","roleName":"Co-Principal Investigator","creator":{"@type":"Person","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Person","@id":"https:\/\/www.bco-dmo.org\/person\/50663","name":"Dr Mark Brzezinski","url":"https:\/\/www.bco-dmo.org\/person\/50663"}},{"@type":"Role","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Participant","roleName":"Student","creator":{"@type":"Person","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Person","@id":"https:\/\/www.bco-dmo.org\/person\/471722","name":"Ms Anna James","url":"https:\/\/www.bco-dmo.org\/person\/471722"}}],"citation":"Passow, U., Brzezinski, M., Carlson, C. (2016) Removal of organic carbon by natural bacterioplankton communities as a function of pCO2 from laboratory experiments between 2012 and 2016. Biological and Chemical Oceanography Data Management Office (BCO-DMO). Dataset version 2013-11-21 [if applicable, indicate subset used]. doi:10.1575\/1912\/bco-dmo.665253 [access date]","version":"2013-11-21","license":"http:\/\/creativecommons.org\/licenses\/by\/4.0\/","publishingPrinciples":{"@type":"DigitalDocument","@id":"http:\/\/creativecommons.org\/licenses\/by\/4.0\/","additionalType":"https:\/\/geodex.org\/voc\/Protocol-License","name":"Dataset Usage License","url":"http:\/\/creativecommons.org\/licenses\/by\/4.0\/"},"temporalCoverage":"2012-09-20\/2016-01-22","spatialCoverage":{"@type":"Place","geo":{"@type":"GeoShape","box":"-17.45,-149.8727 34.407,-64.6353","polygon":"-17.45,-149.8727 34.407,-149.8727 34.407,-64.6353 -17.45,-64.6353 -17.45,-149.8727"},"subjectOf":{"@type":"CreativeWork","fileFormat":"application\/vnd.geo+json","text":"{\u0022type\u0022:\u0022Feature\u0022,\u0022geometry\u0022:{\u0022type\u0022:\u0022Polygon\u0022,\u0022coordinates\u0022:[[[-64.6353,34.407],[-149.8727,34.407],[-149.8727,-17.45],[-64.6353,-17.45],[-64.6353,34.407]]],\u0022properties\u0022:[]}}"},"additionalProperty":[{"@type":"PropertyValue","alternateName":"CRS","name":"Coordinate Reference System","value":"http:\/\/www.opengis.net\/def\/crs\/OGC\/1.3\/CRS84"},{"@type":"PropertyValue","value":"POLYGON ((-149.8727 -17.45, -64.6353 -17.45, -64.6353 34.407, -149.8727 34.407, -149.8727 -17.45))","propertyID":"http:\/\/www.opengis.net\/ont\/geosparql#wktLiteral","name":"Well-Known Text","alternateName":"WKT"}]},"publisher":{"@id":"https:\/\/www.bco-dmo.org","@type":"Organization","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Organization","legalName":"Biological and Chemical Data Management Office","name":"BCO-DMO","identifier":"http:\/\/lod.bco-dmo.org\/id\/affiliation\/191","url":"https:\/\/www.bco-dmo.org","sameAs":"http:\/\/www.re3data.org\/repository\/r3d100000012"},"provider":{"@id":"https:\/\/www.bco-dmo.org"},"includedInDataCatalog":[{"@id":"http:\/\/www.bco-dmo.org\/datasets"}],"recordedAt":[{"@type":"Event","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Cruise","@id":"https:\/\/www.bco-dmo.org\/deployment\/665488","name":"KM1416","description":"","location":{"@type":"Place","name":"South Pacific Subtropical Gyre","address":"South Pacific Subtropical Gyre"},"url":"https:\/\/www.bco-dmo.org\/deployment\/665488","recordedIn":{"@type":"CreativeWork","isBasedOn":{"@type":"Vehicle","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Platform","@id":"https:\/\/www.bco-dmo.org\/platform\/54009","name":"R\/V Kilo Moana","url":"https:\/\/www.bco-dmo.org\/platform\/54009"}},"startDate":"2014-07-19","endDate":"2014-08-08"}],"isPartOf":[{"@type":"CreativeWork","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Project","@id":"https:\/\/www.bco-dmo.org\/project\/2284","name":"Will high CO2 conditions affect production, partitioning and fate of organic matter?","alternateName":"OA - Effects of High CO2","description":"From the NSF Award Abstract\nCoastal waters are already experiencing episodic exposure to carbonate conditions that were not expected until the end of the century making understanding the response to these episodic events as important as understanding the long-term mean response. Among the most striking examples are those associated with coastal upwelling along the west coast of the US, where the pH of surface waters may drop to 7.6 and pCO2 can reach 1100 uatm. Upwelling systems are responsible for a significant fraction of global carbon export making them prime targets for investigations on how ocean acidification is already affecting the biological pump today.\nIn this study, researchers at the University of California at Santa Barbara will investigate the potential effects of ocean acidification on the strength of the biological pump under the transient increases in CO2 experienced due to upwelling. Increases in CO2 are expected to alter the path and processing of carbon through marine food webs thereby strengthening the biological pump. Increases in inorganic carbon without proportional increases in nutrients result in carbon over-consumption by phytoplankton. How carbon over-consumption affects the strength of the biological pump will depend on the fate of the extra carbon that is either incorporated into phytoplankton cells forming particulate organic matter (POM), or is excreted as dissolved organic matter (DOM). Results from mesocosm experiments demonstrate that the mechanisms controlling the partitioning of fixed carbon between the particulate and dissolved phases, and the processing of those materials, are obscured when both processes operate simultaneously under natural or semi-natural conditions. Here, POM and DOM production and the heterotrophic processing of these materials will be separated experimentally across a range of CO2 concentrations by conducting basic laboratory culture experiments. In this way the mechanisms whereby elevated CO2 alters the flow of carbon along these paths can be elucidated and better understood for use in mechanistic forecasting models.\nBroader Impacts- The need to understand the effects of ocean acidification for the future of society is clear. In addition to research education, both formal and informal, will be important for informing the public. Within this project 1-2 graduate students and 2-3 minority students will be recruited as interns from the CAMP program (California Alliance for Minority Participation). Within the \u0027Ocean to Classrooms\u0027 program run by outreach personnel from UCSB\u0027s Marine Science Institute an educational unit for K-12 students will be developed. Advice and support is also given to the Education Coordinator of NOAA, Channel Islands National Marine Sanctuary for the development of an education unit on ocean acidification.\n\nPUBLICATIONS PRODUCED AS A RESULT OF THIS RESEARCH\nArnosti C, Grossart H-P, Muehling M, Joint I, Passow U. \u0026quot;Dynamics of extracellular enzyme activities in seawater under changed atmsopheric pCO2: A mesocosm investigation.,\u0026quot; Aquatic Microbial Ecology, v.64, 2011, p. 285.\nPassow U. \u0026quot;The Abiotic Formation of TEP under Ocean Acidification Scenarios.,\u0026quot; Marine Chemistry, v.128-129, 2011, p. 72.\nPassow, Uta; Carlson, Craig A.. \u0026quot;The biological pump in a high CO2 world,\u0026quot; MARINE ECOLOGY PROGRESS SERIES, v.470, 2012, p. 249-271.\nGaerdes, Astrid; Ramaye, Yannic; Grossart, Hans-Peter; Passow, Uta; Ullrich, Matthias S.. \u0026quot;Effects of Marinobacter adhaerens HP15 on polymer exudation by Thalassiosira weissflogii at different N:P ratios,\u0026quot; MARINE ECOLOGY PROGRESS SERIES, v.461, 2012, p. 1-14.\nPhilip Boyd, Tatiana Rynearson, Evelyn Armstrong, Feixue Fu, Kendra Hayashi, Zhangi Hu, David Hutchins, Raphe Kudela, Elena Litchman, Margaret Mulholland, Uta Passow, Robert Strzepek, Kerry Whittaker, Elizabeth Yu, Mridul Thomas. \u0026quot;Marine Phytoplankton Temperature versus Growth Responses from Polar to Tropical Waters - Outcome of a Scientific Community-Wide Study,\u0026quot; PLOS One 8, v.8, 2013, p. e63091.\nArnosti, C., B. M. Fuchs, R. Amann, and U. Passow. \u0026quot;Contrasting extracellular enzyme activities of particle-associated bacteria from distinct provinces of the North Atlantic Ocean,\u0026quot; Frontiers in Microbiology, v.3, 2012, p. 1.\nKoch, B.P., Kattner, G., Witt, M., Passow, U., 2014. Molecular insights into the microbial formation of marine dissolved organic matter: recalcitrant or labile? Biogeosciences Discuss. 11 (2), 3065-3111.\nTaucher, J., Brzezinski, M., Carlson, C., James, A., Jones, J., Passow, U., Riebesell, U., submitted. Effects of warming and elevated pCO2 on carbon uptake and partitioning of the marine diatoms Thalassiosira weissflogii and Dactyliosolen fragilissimus. Limnology and Oceanography\n","url":"https:\/\/www.bco-dmo.org\/project\/2284"}],"distribution":[{"@type":"DataDownload","@id":"https:\/\/www.bco-dmo.org\/dataset\/472032\/data\/download","contentUrl":"https:\/\/www.bco-dmo.org\/dataset\/472032\/data\/download","encodingFormat":"text\/tab-separated-values","datePublished":"2016-12-05","inLanguage":"en-US"}],"measurementTechnique":["Shimadzu TOC-V Analyzer","Flow Cytometer","Microscope-Fluorescence","Shimadzu TOC-L Analyzer"],"variableMeasured":[{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665785","name":"experiment","description":"Experiment identifier\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665785","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665786","name":"site","description":"Site the water for the experiment came from\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665786","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665787","name":"latitude","description":"Latitude where water samples were collected; north is positive.\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665787","unitText":"decimal degrees"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665788","name":"longitude","description":"Longitude where water samples were collected; west is negative.\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665788","unitText":"decimal degrees"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665789","name":"bottle_number","description":"Bottle identifier\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665789","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665790","name":"doc_addition","description":"Dissolved organic carbon additions. See Aquisition Description section for an explaination\u00a0of values.\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665790","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665791","name":"target_pCO2","description":"Target pCO2 level\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665791","unitText":"parts per million (ppm)"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665792","name":"time_point","description":"Time point identifier in experiment\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665792","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665793","name":"time_days","description":"Elapsed time since start of experiment in days\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665793","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665794","name":"date","description":"Date of experiment in format YYYY-MM-DD\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665794","unitText":"unitless"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665795","name":"bact_abun_x10e6_avg","description":"Bacterial abundance multiplied by 10^6\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665795","unitText":"cells per milliliter"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665796","name":"bact_abun_x10e6_stderr","description":"Standard error of bacterial abundance multiplied by 10^6\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665796","unitText":"cells per milliliter"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665797","name":"bact_abun_x10e6_stdev","description":"Standard deviation Bacterial abundance multiplied by 10^6\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665797","unitText":"cells per milliliter"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665798","name":"toc_avg","description":"Total organic carbon\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665798","unitText":"micromoles per liter (uM)"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665799","name":"toc_stderr","description":"Standard error of total organic carbon\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665799","unitText":"micromoles per liter (uM)"},{"@type":"PropertyValue","additionalType":"https:\/\/library.ucar.edu\/earthcollab\/schema#Parameter","@id":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665800","name":"toc_stdev","description":"Standard deviation of total organic carbon\n","url":"https:\/\/www.bco-dmo.org\/dataset-parameter\/665800","unitText":"micromoles per liter (uM)"}],"funder":[{"@type":"Organization","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Organization","legalName":"NSF Division of Ocean Sciences","name":"NSF OCE","makesOffer":[{"@id":"https:\/\/www.bco-dmo.org\/award\/55209","@type":"Offer","additionalType":"http:\/\/schema.geolink.org\/1.0\/base\/main#Award","name":"OCE-1041038","url":"https:\/\/www.bco-dmo.org\/award\/55209","sameAs":"http:\/\/www.nsf.gov\/awardsearch\/showAward?AWD_ID=1041038\u0026HistoricalAwards=false","offeredBy":{"@type":"Person","@id":"https:\/\/www.bco-dmo.org\/person\/51467","additionalType":"http:\/\/schema.geolink.org\/1.0\/voc\/local#roletype_program_manager","name":"Dr Donald  L. Rice","url":"https:\/\/www.bco-dmo.org\/person\/51467"}}]}],"gdx:fundedBy":[{"@id":"https:\/\/www.bco-dmo.org\/award\/55209"}]}`

const iedaUrl string = "http://get.iedadata.org/doi/100700"
const iedaJsonld string = `{
  "@context": {
 "@vocab": "http://schema.org/",
"datacite": "http://purl.org/spar/datacite/",
                "earthcollab": "https://library.ucar.edu/earthcollab/schema#",
                "geolink": "http://schema.geolink.org/1.0/base/main#",
                "vivo": "http://vivoweb.org/ontology/core#",
                "dcat":"http://www.w3.org/ns/dcat#"
                  },
 "@id": "DOI:10.1594/IEDA/100700",
  "@type": "Dataset",
  "additionalType": [
    "http://schema.geolink.org/1.0/base/main#Dataset",
    "http://vivoweb.org/ontology/core#Dataset"
  ],
  "name": "S, Se and Te contents of basalts from the Reykjanes Ridge and SW Iceland Rift Zone",
  "citation": "Allison Forrest, Jean-Guy Schilling, Katherine A. Kelley (2017), S, Se and Te contents of basalts from the Reykjanes Ridge and SW Iceland Rift Zone. Interdisciplinary Earth Data Alliance (IEDA). doi:10.1594/IEDA/100700",
  "creator":[{
                "@type": "Person",
                "additionalType": "http://schema.geolink.org/1.0/base/main#Person",
      "name": "Allison Forrest",
      "givenName": "Allison",
      "familyName": "Forrest"},
{
                "@type": "Person",
                "additionalType": "http://schema.geolink.org/1.0/base/main#Person",
      "name": "Jean-Guy Schilling",
      "givenName": "Jean-Guy",
      "familyName": "Schilling"},
{
                "@type": "Person",
                "additionalType": "http://schema.geolink.org/1.0/base/main#Person",
      "name": "Katherine A. Kelley",
      "givenName": "Katherine",
      "familyName": "Kelley"}],
  "datePublished": "2017",
  "dateCreated": "2017-06-28",
  "version": "1.0",
  "inLanguage": "en",
  "description": "Abstract: This data set reports S, Se, and Te concentrations in basaltic whole-rocks and glasses from the SW Iceland rift zone, including sub-glacial glasses, and submarine lavas from the Reykjanes Ridge, Gibbs Fracture Zone, and Mid-Atlantic Ridge. We thank R. Kingsley, J.M. Rhodes, and C. Mandeville for assistance with analyses. This work was supported by NSF award OCE# 0326658 to JGS. NSF award OCE# 1555523 provides curatorial support for marine geological samples at the Graduate School of Oceanography, University of Rhode Island.; Other Description: Forrest, A. (2005), Sulfur, Selenium and Tellurium in Reykjanes Ridge and Iceland Basalts, 175 pp, University of Rhode Island, Kingston, RI. Forrest, A., R. Kingsley, and J.-G. Schilling (2009), Determination of Selenium and Tellurium in Basalt Rock Reference Materials by Isotope Dilution Hydride Generation-Inductively Coupled Plasma-Mass Spectrometry (ID-HG-ICP-MS), Geostandards and Geoanalytical Research, 33(2), 261-269, doi:10.1111/j.1751-908X.2009.00841.x.",
  "distribution": [
{
    "@type": "DataDownload",
    "additionalType": "http://www.w3.org/ns/dcat#distribution",
      "name":"DOI landing page",
      "http://www.w3.org/ns/dcat#accessURL": "http://dx.doi.org/10.1594/IEDA/100700",
      "url": "http://dx.doi.org/10.1594/IEDA/100700",
      "encodingFormat": ["application/pdf", "application/vnd.ms-excel"]},
{
    "@type": "DataDownload",
    "additionalType": "http://www.w3.org/ns/dcat#distribution",
      "name":"URL",
      "http://www.w3.org/ns/dcat#accessURL": "http://www.earthchem.org/library/browse/view?id=1085",
      "url": "http://www.earthchem.org/library/browse/view?id=1085"
,
      "encodingFormat": ["application/pdf", "application/vnd.ms-excel"]}  ],
  "identifier": [            {
                "@id": "doi:10.1594/IEDA/100700",
               "@type": "PropertyValue",
            "additionalType": ["http://schema.geolink.org/1.0/base/main#Identifier", "http://purl.org/spar/datacite/Identifier"],
            "propertyID": "http://purl.org/spar/datacite/doi",
      "url": "http://dx.doi.org/10.1594/IEDA/100700",      "value": "10.1594/IEDA/100700"}  ],
  "isPartOf": [{
     "@id":"http://www.nsf.gov/awardsearch/showAward.do?AwardNumber=0326658",
     "@type": "CreativeWork",
                "additionalType": "https://library.ucar.edu/earthcollab/schema#Project",
     "funder": {
       "@type":"Organization",
       "name":"National Science Foundation"       }
     },
{
     "@id":"http://www.nsf.gov/awardsearch/showAward.do?AwardNumber=1555523",
     "@type": "CreativeWork",
                "additionalType": "https://library.ucar.edu/earthcollab/schema#Project",
     "funder": {
       "@type":"Organization",
       "name":"National Science Foundation"       }
     }],
  "keywords": ["Regional (Continents, Oceans)", "Spreading Center", "Chemistry:Rock", "Geochemistry", "Marine Geoscience", "Solid Earth", "basalt", "mid-ocean ridge", "plume", "Selenium", "sulfur", "Tellurium", "Gibbs Fracture Zone", "Iceland", "Mid-Atlantic Ridge", "Reykjanes Ridge",
"Reykjanes Ridge, Iceland, Mid-Atlantic Ridge, Gibbs Fracture Zone"],
  "license": "Creative Commons Attribution-NonCommercial-Share Alike 3.0 United States [CC BY-NC-SA 3.0]",
  "provider":
           {
    "@type": "Organization",
    "@id": "https://www.iedadata.org/",
    "name": "Interdisciplinary Earth Data Alliance (IEDA)"
            },
  "publisher": {

    "@type": "Organization",
    "@id": "https://www.iedadata.org/",
    "name": "Interdisciplinary Earth Data Alliance (IEDA)",
    "url": "https://www.iedadata.org/",
    "description": "The IEDA data facility mission is to support, sustain, and advance the geosciences by providing data services for observational geoscience data from the Ocean, Earth, and Polar Sciences. IEDA systems serve as primary community data collections for global geochemistry and marine geoscience research and support the preservation, discovery, retrieval, and analysis of a wide range of observational field and analytical data types. Our tools and services are designed to facilitate data discovery and reuse for focused disciplinary research and to support interdisciplinary research and data integration.",
    "logo": {
      "@type": "ImageObject",
      "url": "http://app.iedadata.org/images/ieda_maplogo.png"
    },
    "contactPoint": {
      "@type": "ContactPoint",
      "name": "Information Desk",
      "email": "info@iedadata.org",
      "url": "https://www.iedadata.org/contact/",
      "contactType": "Information"
    },
    "parentOrganization": {
      "@type": "Organization",
      "@id": "https://viaf.org/viaf/142992181/",
      "name": "Lamont-Doherty Earth Observatory",
      "url": "http://www.ldeo.columbia.edu/",
      "address": {
        "@type": "PostalAddress",
        "streetAddress": "61 Route 9W",
        "addressLocality": "Palisades",
        "addressRegion": "NY",
        "postalCode": "10964-1000",
        "addressCountry": "USA"
      },
      "parentOrganization": {
        "@type": "Organization",
        "@id": "https://viaf.org/viaf/156836332/",
        "legalName": "Columbia University",
        "url": "http://www.columbia.edu/"
      }
            }
        ,
            "funder":
          {
            "@type": "Organization",
            "@id": "http://dx.doi.org/10.13039/100000085",
            "legalName": "Directorate for Geosciences",
            "alternateName": "NSF-GEO",
            "url": "http://www.nsf.gov",
            "parentOrganization": {
                "@type": "Organization",
                "@id": "http://dx.doi.org/10.13039/100000001",
                "legalName": "National Science Foundation",
                "alternateName": "NSF",
                "url": "http://www.nsf.gov"
            }
           }
            ,
            "publishingPrinciples": {

      "@id": "http://creativecommons.org/licenses/by-nc-sa/3.0/us/",
      "@type": "DigitalDocument",
      "additionalType": "gdx:Protocol-License",
      "name": "Dataset Usage License",
      "description": "Creative Commons Attribution-NonCommercial-Share Alike 3.0 United States [CC BY-NC-SA 3.0]",
      "url": "https://creativecommons.org/licenses/by-nc-sa/3.0/us/"

       }
                },
  "spatialCoverage": [{
"@type": "Place",
      "name": "Reykjanes Ridge, Iceland, Mid-Atlantic Ridge, Gibbs Fracture Zone",
      "geo":[{
      "@type": "GeoShape",
      "box": "-35.4, 50.46 -17.68, 65"},
      {
      "@type": "GeoCoordinates",
      "longitude":"-29.42",
      "latitude":"50.46"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-30.02",
      "latitude":"51.28"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.92",
      "latitude":"51.56"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.95",
      "latitude":"52.01"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-31.52",
      "latitude":"52.33"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-31.57",
      "latitude":"52.48"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-34.94",
      "latitude":"52.66"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-35.24",
      "latitude":"53.41"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-35.4",
      "latitude":"54.25"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-35.22",
      "latitude":"54.76"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-34.87",
      "latitude":"55.67"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-33.42",
      "latitude":"57.14"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-32.57",
      "latitude":"57.68"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-31.62",
      "latitude":"58.42"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-30.95",
      "latitude":"58.87"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.53",
      "latitude":"59.99"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.51",
      "latitude":"59.99"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.43",
      "latitude":"59.99"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.48",
      "latitude":"60.02"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-29.38",
      "latitude":"60.03"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-28.88",
      "latitude":"60.45"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-28.42",
      "latitude":"60.73"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-27.9",
      "latitude":"61.09"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-27.88",
      "latitude":"61.1"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-27.07",
      "latitude":"61.6"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-26.88",
      "latitude":"61.73"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-26.54",
      "latitude":"61.98"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-26.29",
      "latitude":"62.08"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-26.36",
      "latitude":"62.11"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-26.14",
      "latitude":"62.26"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.93",
      "latitude":"62.31"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.78",
      "latitude":"62.35"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.84",
      "latitude":"62.37"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.45",
      "latitude":"62.59"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.43",
      "latitude":"62.62"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.21",
      "latitude":"62.7"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-25.16",
      "latitude":"62.79"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.88",
      "latitude":"62.9"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.69",
      "latitude":"62.99"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.5",
      "latitude":"63.07"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.46",
      "latitude":"63.19"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.27",
      "latitude":"63.22"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.2",
      "latitude":"63.27"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-24.243",
      "latitude":"63.277"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-23.85",
      "latitude":"63.46"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-23.87",
      "latitude":"63.47"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-23.8",
      "latitude":"63.47"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-23.7",
      "latitude":"63.57"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-22.7",
      "latitude":"63.85"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-22.43",
      "latitude":"63.86"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-22.05",
      "latitude":"63.9"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-22.54",
      "latitude":"63.91"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-22.2",
      "latitude":"63.95"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-21.74",
      "latitude":"63.97"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-22.57",
      "latitude":"63.98"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-21.97",
      "latitude":"63.98"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-21.9",
      "latitude":"64.1"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-18.99",
      "latitude":"64.23"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-20.435",
      "latitude":"64.277"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-21.11",
      "latitude":"64.28"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-21",
      "latitude":"64.33"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-21.08",
      "latitude":"64.45"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-19.155",
      "latitude":"64.663"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-17.68",
      "latitude":"64.68"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-20.78",
      "latitude":"64.73"}]},
{
"@type": "Place",
      "geo":[      {
      "@type": "GeoCoordinates",
      "longitude":"-19.58",
      "latitude":"65"}]}]}`
