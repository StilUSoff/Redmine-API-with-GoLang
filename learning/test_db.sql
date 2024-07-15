CREATE TABLE ships (
    class VARCHAR(1000),
    name VARCHAR(1000),
    country VARCHAR(1000),
    organization VARCHAR(1000),
    displacement VARCHAR(1000),
    length VARCHAR(1000),
    width VARCHAR(1000),
    draft VARCHAR(1000),
    propulsion VARCHAR(1000),
    engines VARCHAR(1000),
    power VARCHAR(1000),
    speed VARCHAR(1000),
    range VARCHAR(1000),
    autonomy VARCHAR(1000),
    crew VARCHAR(1000),
    artillery VARCHAR(1000),
    anti_aircraft VARCHAR(1000),
    tactical_strike VARCHAR(1000),
    navigation VARCHAR(1000),
    radar VARCHAR(1000),
    electronic_warfare VARCHAR(1000),
    missile VARCHAR(1000),
    anti_submarine VARCHAR(1000),
    mine_torpedo VARCHAR(1000),
    aviation_group VARCHAR(1000),
    model_path VARCHAR(1000)
);

INSERT INTO ships (class, name, country, organization, displacement, length, width, draft, propulsion, engines, power, speed, range, autonomy, crew, artillery, anti_aircraft, tactical_strike, navigation, radar, electronic_warfare, missile, anti_submarine, mine_torpedo, aviation_group, model_path) 
VALUES 
('destroyer', 'HMS Daring', 'United Kingdom', 'Royal Navy', '5893 (standard), 7570 (full), 8500', '152.4 m', '21.2 m', '5.3 m', '2 fixed-pitch propellers', '2 x Rolls-Royce WR-21 gas turbines, 2 x Wartsila V12 VASA32 diesel generators, 2 x Converteam electric motors', '2 x 28,800 hp, 2 x 2700 hp, 2 x 27,000 hp', '31 knots', '6500 miles at 18 knots', '45 days', '191 (+41 extra spaces)', '1 x 114 mm Mark 8 Mod. 1', '2 x 20 mm Mark 15 Phalanx CIWS', '', 'multi-function radar Sampson', 'S1850 long-range air and surface search radar', '', '', 'MFS-7000 sonar', 'anti-torpedo defense system', 'hangar, 1 Lynx HMA8 or Merlin HM1 helicopter', 'C:\\Models\\ship1.stl'),
('littoral combat ship', 'LCS', 'USA', 'US Navy', '2839 tons', '115.3 m', '17.5 m', '3.7—4.1 m', '4 waterjet propulsors', 'gas turbine and diesel engine', 'Turbines 2 x 29,500 hp, diesels 2 x 12,203 hp', '50 knots', '4300 miles at 20 knots, 1500 miles at 50 knots', '21 days', '50 people', '1 x 1 — 57 mm Mk. 110', '4 x 12.7 mm machine guns', '', '', 'SAAB Sea Giraffe AMB radar', '', '1 x 21 RAM Mk. 31 SAM launcher', 'Honeywell Mark 50', '', '2 SH-60 helicopters or 1 H-60 helicopter and 3 MQ-8 Fire Scout UAVs', 'C:\\Models\\ship2.stl'),
('battleship', 'Iowa', 'USA', 'US Navy', '48,800 (standard), 58,460 (full)', '270.5 m', '33 m', '11 m', '4 x propellers; 4 x reduction gear steam turbines', '4 shaft steam turbines', '212,000 hp', '33-35 knots', '14,890 miles at 15 knots', '', '1800 people', '9 (3x3) — 406 mm/50 Mk.7, 10 x 2 — 127 mm/38 Mk.12', '4 Vulcan Phalanx CIWS', '32 Tomahawk missiles', '', '', '', '4 x 4 Harpoon missile launchers', '', '2 triple 324 mm torpedo tubes', '2 helicopters', 'C:\\Models\\ship3.stl'),
('frigate', 'Iver Huitfeldt', 'Denmark', 'Royal Danish Navy', '5850 (standard), 6645 (full)', '138.7 m', '19.8 m', '5.3 m', '2 fixed-pitch propellers', '4 MTU 8000 20V M70 diesels', '21,500 hp', '<26 knots', '9000 miles at 15 knots', '28 days', '100 people', '2 x 76 mm OTO Breda', '35 mm Oerlikon Millennium', '', '', 'SMART-L radar, APAR CEROS 200 radar, ES-3701 EW system', '', '', '', '2 x 2 MU90 torpedo tubes', '1 EH-101 helicopter', 'C:\\Models\\ship4.stl'),
('frigate', 'Oliver Hazard Perry', 'USA — 29, Turkey — 8, Republic of China — 8, Australia — 6, Spain — 6, Egypt — 4, Poland — 2, Bahrain — 1, Pakistan — 1', 'Foreign navies', '4200', '138.1 m', '13.7 m', '6.7—7.77 m', '', '2 General Electric LM2500-30 gas turbines, 2 auxiliary engines', '41,000 hp', '29 knots', '4500 miles at 20 knots, 5000 miles at 18 knots', '', '219 people (including 19 officers)', '1 x 1 — 76 mm/62 Compact OTO Melara', '1 x 20 mm Phalanx CIWS', '', '', 'Radars: AN/SPS-49, AN/SPS-55, CAS, STIR, Sonar: AN/SQS-56, AN/SQR-19', 'NTDS, AN/SQQ-89, Mk 92 FCS, LAMPS, SLQ-32(V)2 EW system', '4 RGM-84 Harpoon missiles in launchers, 1 Mk 13 SAM launcher (36 SM-1 missiles)', '', '2 triple Mark 32 ASW torpedo tubes (6 Mark 46 or Mark 50 torpedoes)', '1 SH-2 Seasprite helicopter (on short ships) or 2 SH-60 Seahawk helicopters (on others)', 'C:\\Models\\ship5.stl'),
('amphibious assault ship', 'Tarawa', 'USA', 'US Navy', '33,536 (standard), 39,967 (full)', '254.20 m', '40.23 m', '9.72 m', '2 x shafts, 1 x bow thruster', 'steam plant, 2 x fuel-burning boilers, 2 x Westinghouse turbines', '70,000 hp', '24 knots', '10,000 miles at 20 knots', '', '500 marines', '', '', '', '', '', '', 'Mark 49 RAM missile system, 2 x Phalanx CIWS, 6 x 25 mm automatic cannons, 8 x 12.7 mm machine guns', '', '', 'Flight deck size 820 x 118.1 feet (249.9 x 36.0 m) with 2 aircraft elevators', 'C:\\Models\\ship6.stl'),
('guided missile cruiser', 'Ticonderoga', 'USA', 'US Navy', '9800', '172.8 m', '16.8 m', '10.2 m', '', '4 General Electric LM2500 gas turbines', '80,000 hp', '32.5 knots', '6000 miles at 20 knots', '', '387 people', '2 x Mk 45 127 mm', '2 x 20 mm Phalanx CIWS; 2 x 25 mm Mk 38', '', '', '', '', '26 Tomahawk missiles, 16 ASROC missiles, 80 Standard 2 SAMs. Total armament: 122 missiles', '', '2 triple 324 mm torpedo tubes', '2 Sikorsky SH-60B or MH-60R Seahawk helicopters', 'C:\\Models\\ship7.stl')