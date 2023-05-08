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

import { ModuleDetailProps } from './Model'
import ModuleDownload from './ModuleDownload';

const ModuleDetail = (props: ModuleDetailProps) => {
  return(
    <>
      <div id="tab1">
      Module Name: {props.module.name}
      Module Description: {props.module.description}
      </div>

      <div id="tab2">
        Parse the markdown into HTML, ensure Mermaid is supported
        {props.module.detailMarkdown}
      </div>

      <div>
        
        <ModuleDownload />
      </div>

    </>
  )
}

export default ModuleDetail;