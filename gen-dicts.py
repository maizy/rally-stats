#!/usr/bin/env python3
import json
import math
import os.path
import sys
import sqlite3
import collections

PROJECT_DIR = os.path.dirname(__file__)
DICTS_DIR = os.path.join(PROJECT_DIR, 'rstats_app', 'dicts')

LOCATIONS = {
    100: 'Hawkes Bay, New Zealand',
    200: 'Värmland, Sweden',
    300: 'Powys, Wales',
    400: 'Leczna County, Poland',
    500: 'Ribadelles, Spain',
    600: 'Monaro, Australia',
    700: 'Baumholder, Germany',
    800: 'Catamarca, Argentina',
    900: 'New England, USA',
    1000: 'Monte Carlo, Monaco',
    1100: 'Argolis, Greece',
    1200: 'Jämsä, Finland',
    1300: 'Perth and Kinross, Scotland',
}

CarClass = collections.namedtuple('CarClass', ['name', 'drivetrain'])

CAR_CLASSES = {
    100: CarClass('Historic Rally H1 (FWD)', 'FWD'),
    200: CarClass('Historic Rally H2 (FWD)', 'FWD'),
    300: CarClass('Historic Rally H2 (RWD)', 'RWD'),
    400: CarClass('Historic Rally H3 (RWD)', 'RWD'),
    500: CarClass('Historic Rally Group B (RWD)', 'RWD'),
    600: CarClass('Historic Rally Group B (4WD)', '4WD'),
    700: CarClass('Modern Rally R2', 'FWD'),
    800: CarClass('Modern Rally Group A', '4WD'),
    900: CarClass('Modern Rally NR4/R4', '4WD'),
    1000: CarClass('Up to 2000cc (4WD)', '4WD'),
    1100: CarClass('Modern Rally R5', '4WD'),
    1200: CarClass('Modern Rally GT', 'RWD'),
    1300: CarClass('F2 Kit Car', 'FWD'),
}

if __name__ == '__main__':
    if len(sys.argv) != 2:
        # https://github.com/soong-construction/dirt-rally-time-recorder/blob/master/resources/setup-dr2.sql
        print('usage gen-dicts.py setup-dr2.sql')
        exit(1)

    setup_sql = open(sys.argv[1], 'rb').read().decode('utf-8')
    con = sqlite3.connect(":memory:")
    con.row_factory = sqlite3.Row
    cur = con.cursor()
    cur.executescript(setup_sql)

    # import pprint; pprint.pprint([dict(i) for i in cur.execute("select * from sqlite_schema").fetchall()])

    locations_data = collections.OrderedDict()
    for id, name in LOCATIONS.items():
        locations_data[id] = {'name': name}
    json.dump(locations_data, fp=open(os.path.join(DICTS_DIR, 'locations.json'), 'w'), ensure_ascii=False, indent=2)

    tracks = collections.OrderedDict()
    for track in cur.execute("select * from Tracks order by id").fetchall():
        location_id = int(math.floor(track['id'] / 100.0)) * 100
        if location_id not in LOCATIONS:
            print(f'Unable to detect location for track {dict(track)}')
            exit(2)
        tracks[track['id']] = {
            'name': track['name'],
            'length': track['length'],
            'location_id': location_id,
        }
    json.dump(tracks, fp=open(os.path.join(DICTS_DIR, 'tracks.json'), 'w'), ensure_ascii=False, indent=2)

    car_classes_data = collections.OrderedDict()
    for id, spec in CAR_CLASSES.items():
        car_classes_data[id] = spec._asdict()
    json.dump(car_classes_data, fp=open(os.path.join(DICTS_DIR, 'car_classes.json'), 'w'), ensure_ascii=False, indent=2)

    cars = collections.OrderedDict()
    for car in cur.execute('select * from cars').fetchall():
        class_id = int(math.floor(car['id'] / 100.0)) * 100
        if class_id not in CAR_CLASSES:
            print(f'Unable to detect car class for car {dict(car)}')
            exit(2)
        cars[car['id']] = {
            'name': car['name'],
            'car_class_id': class_id,
        }
    json.dump(cars, fp=open(os.path.join(DICTS_DIR, 'cars.json'), 'w'), ensure_ascii=False, indent=2)

    print('done')
