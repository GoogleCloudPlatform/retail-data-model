// Copyright 2023 Ryan McGuinness
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { RDMModule } from "./Model";

export const ModuleData: Array<RDMModule> = [
  {
    name: "Common",
    version: "1.0.0",
    banner: 'assets/images/rdm-common.png',
    bannerAlt: 'Retail Data Model Common banner',
    description: "The common data model...",
    detailMarkdown: "",
    features: [],
    terraformProperties: []
  },
  {
    name: "Events",
    version: "1.0.0",
    banner: 'assets/images/rdm-common.png',
    bannerAlt: 'Retail Data Model Events',
    description: "The common events ",
    detailMarkdown: "",
    features: [],
    terraformProperties: []
  },
];